/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/ptr"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/prasad89/devspace-operator/api/v1alpha1"
)

// DevSpaceReconciler reconciles a DevSpace object
type DevSpaceReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// +kubebuilder:rbac:groups=platform.devspace.io,resources=devspaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=platform.devspace.io,resources=devspaces/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=platform.devspace.io,resources=devspaces/finalizers,verbs=update

// Standard K8s resources we will create/manage
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services;persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete

func (r *DevSpaceReconciler) addPvcIfNotExist(ctx context.Context, devspace *v1alpha1.DevSpace) error {
	log := log.FromContext(ctx)

	pvc := &corev1.PersistentVolumeClaim{}
	pvcName := devspace.Name
	namespace := devspace.Namespace

	// Try fetching the PVC
	err := r.Get(ctx, client.ObjectKey{Namespace: namespace, Name: pvcName}, pvc)
	if err == nil {
		log.Info("PVC already exists", "pvc", pvcName)
		return nil // PVC exists
	}
	if !apierrors.IsNotFound(err) {
		log.Error(err, "Failed to check if PVC exists")
		return err
	}

	// PVC not found — create it
	desiredPVC := generateDesiredPVC(devspace, pvcName)
	if err := controllerutil.SetControllerReference(devspace, desiredPVC, r.Scheme); err != nil {
		log.Error(err, "Failed to set owner reference")
		return err
	}

	if err := r.Create(ctx, desiredPVC); err != nil {
		log.Error(err, "Failed to create PVC")
		return err
	}

	if r.Recorder != nil {
		r.Recorder.Event(devspace, corev1.EventTypeNormal, "PVCReady", "PVC created successfully")
	}
	log.Info("PVC created", "pvc", pvcName)
	return nil
}

func generateDesiredPVC(devspace *v1alpha1.DevSpace, pvcName string) *corev1.PersistentVolumeClaim {
	return &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pvcName,
			Namespace: devspace.Namespace,
			Labels: map[string]string{
				"app":      "devspace",
				"devspace": devspace.Name,
			},
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse("1Gi"), // Hardcoded value
				},
			},
		},
	}
}

func (r *DevSpaceReconciler) addOrUpdateSts(ctx context.Context, devspace *v1alpha1.DevSpace) error {
	log := log.FromContext(ctx)

	stsList := &appsv1.StatefulSetList{}
	labelSelector := labels.Set{"app": devspace.Name}.AsSelector()
	namespace := devspace.Namespace
	stsName := devspace.Name

	// Try fetching the existing StatefulSet(s)
	err := r.List(ctx, stsList, &client.ListOptions{
		Namespace:     namespace,
		LabelSelector: labelSelector,
	})
	if err != nil {
		log.Error(err, "Failed to list StatefulSets")
		return err
	}

	if len(stsList.Items) > 0 {
		log.Info("StatefulSet already exists", "name", stsName)
		return nil // StatefulSet exists — skipping creation
	}

	// StatefulSet not found — create it
	desiredSts := generateDesiredSts(devspace)
	if err := controllerutil.SetControllerReference(devspace, desiredSts, r.Scheme); err != nil {
		log.Error(err, "Failed to set owner reference on StatefulSet")
		return err
	}

	if err := r.Create(ctx, desiredSts); err != nil {
		log.Error(err, "Failed to create StatefulSet")
		return err
	}

	if r.Recorder != nil {
		r.Recorder.Event(devspace, corev1.EventTypeNormal, "STSCreated", "StatefulSet created successfully")
	}

	log.Info("StatefulSet created", "sts", stsName)
	return nil
}

func generateDesiredSts(devspace *v1alpha1.DevSpace) *appsv1.StatefulSet {
	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devspace.Name,
			Namespace: devspace.Namespace,
			Labels: map[string]string{
				"app": devspace.Name,
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: devspace.Name,
			Replicas:    ptr.To[int32](1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": devspace.Name,
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": devspace.Name,
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "vscode",
							Image:           "gitpod/openvscode-server",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{ContainerPort: 3000},
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "homedir",
									MountPath: fmt.Sprintf("/home/%s/%s", devspace.Spec.Owner, devspace.Name),
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "homedir",
							VolumeSource: corev1.VolumeSource{
								PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
									ClaimName: devspace.Name,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *DevSpaceReconciler) addOrUpdateSvc(ctx context.Context, devspace *v1alpha1.DevSpace) error {
	log := log.FromContext(ctx)

	svc := &corev1.Service{}
	svcName := devspace.Name
	namespace := devspace.Namespace

	// Try fetching the Service
	err := r.Get(ctx, client.ObjectKey{Namespace: namespace, Name: svcName}, svc)
	if err == nil {
		log.Info("Service already exists", "svc", svcName)
		return nil // Service exists
	}
	if !apierrors.IsNotFound(err) {
		log.Error(err, "Failed to check if Service exists")
		return err
	}

	// Service not found — create it
	desiredSvc := generateDesiredSvc(devspace)
	if err := controllerutil.SetControllerReference(devspace, desiredSvc, r.Scheme); err != nil {
		log.Error(err, "Failed to set owner reference on Service")
		return err
	}

	if err := r.Create(ctx, desiredSvc); err != nil {
		log.Error(err, "Failed to create Service")
		return err
	}

	if r.Recorder != nil {
		r.Recorder.Event(devspace, corev1.EventTypeNormal, "SVCCreated", "Service created successfully")
	}
	log.Info("Service created", "svc", svcName)
	return nil
}

func generateDesiredSvc(devspace *v1alpha1.DevSpace) *corev1.Service {
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devspace.Name,
			Namespace: devspace.Namespace,
			Labels: map[string]string{
				"app": devspace.Name,
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": devspace.Name,
			},
			Ports: []corev1.ServicePort{
				{
					Port:       80,
					TargetPort: intstr.FromInt(3000),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Type: corev1.ServiceTypeNodePort,
		},
	}
}

func (r *DevSpaceReconciler) addOrUpdateIngress(ctx context.Context, devspace *v1alpha1.DevSpace) error {
	log := log.FromContext(ctx)

	ingress := &networkingv1.Ingress{}
	ingressName := devspace.Name
	namespace := devspace.Namespace

	// Try fetching the Ingress
	err := r.Get(ctx, client.ObjectKey{Namespace: namespace, Name: ingressName}, ingress)
	if err == nil {
		log.Info("Ingress already exists", "ingress", ingressName)
		return nil // Ingress exists
	}
	if !apierrors.IsNotFound(err) {
		log.Error(err, "Failed to check if Ingress exists")
		return err
	}

	// Ingress not found — create it
	desiredIngress := generateDesiredIngress(devspace)
	if err := controllerutil.SetControllerReference(devspace, desiredIngress, r.Scheme); err != nil {
		log.Error(err, "Failed to set owner reference on Ingress")
		return err
	}

	if err := r.Create(ctx, desiredIngress); err != nil {
		log.Error(err, "Failed to create Ingress")
		return err
	}

	if r.Recorder != nil {
		r.Recorder.Event(devspace, corev1.EventTypeNormal, "IngressCreated", "Ingress created successfully")
	}
	log.Info("Ingress created", "ingress", ingressName)
	return nil
}

func generateDesiredIngress(devspace *v1alpha1.DevSpace) *networkingv1.Ingress {
	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      devspace.Name,
			Namespace: devspace.Namespace,
			Labels: map[string]string{
				"app": devspace.Name,
			},
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: devspace.Spec.Hostname,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: ptr.To(networkingv1.PathTypePrefix),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: devspace.Name,
											Port: networkingv1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func addCondition(status *v1alpha1.DevSpaceStatus, condType string, statusType metav1.ConditionStatus, reason, message string) {
	for i, existingCondition := range status.Conditions {
		if existingCondition.Type == condType {
			status.Conditions[i].Status = statusType
			status.Conditions[i].Reason = reason
			status.Conditions[i].Message = message
			status.Conditions[i].LastTransitionTime = metav1.Now()
			return
		}
	}

	condition := metav1.Condition{
		Type:               condType,
		Status:             statusType,
		Reason:             reason,
		Message:            message,
		LastTransitionTime: metav1.Now(),
	}
	status.Conditions = append(status.Conditions, condition)
}

func (r *DevSpaceReconciler) updateStatus(ctx context.Context, devspace *v1alpha1.DevSpace) error {
	log := log.FromContext(ctx)

	if err := r.Status().Update(ctx, devspace); err != nil {
		log.Error(err, "Failed to update DevSpace status")
		return err
	}
	log.Info("DevSpace status updated", "status", devspace.Status)
	return nil
}

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the DevSpace object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *DevSpaceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	devspace := &v1alpha1.DevSpace{}
	if err := r.Get(ctx, req.NamespacedName, devspace); err != nil {
		log.Error(err, "Failed to get DevSpace")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	pvcReady, stsReady, svcReady, ingressReady := false, false, false, false

	log.Info("Reconciling DevSpace", "name", devspace.Name, "namespace", devspace.Namespace)

	if err := r.addPvcIfNotExist(ctx, devspace); err != nil {
		log.Error(err, "Failed to add PVC")
		addCondition(&devspace.Status, "PVCNotReady", metav1.ConditionFalse, "PVCNotReady", "Failed to add PVC for DevSpace")
		return ctrl.Result{}, err
	} else {
		pvcReady = true
	}

	if err := r.addOrUpdateSts(ctx, devspace); err != nil {
		log.Error(err, "Failed to add or update StatefulSet")
		addCondition(&devspace.Status, "STSNotReady", metav1.ConditionFalse, "STSNotReady", "Failed to add or update StatefulSet for DevSpace")
		return ctrl.Result{}, err
	} else {
		stsReady = true
	}

	if err := r.addOrUpdateSvc(ctx, devspace); err != nil {
		log.Error(err, "Failed to add or update Service")
		addCondition(&devspace.Status, "ServiceNotReady", metav1.ConditionFalse, "ServiceNotReady", "Failed to add or update Service for DevSpace")
		return ctrl.Result{}, err
	} else {
		svcReady = true
	}

	if err := r.addOrUpdateIngress(ctx, devspace); err != nil {
		log.Error(err, "Failed to add or update Ingress")
		addCondition(&devspace.Status, "IngressNotReady", metav1.ConditionFalse, "IngressNotReady", "Failed to add or update Ingress for DevSpace")
		return ctrl.Result{}, err
	} else {
		ingressReady = true
	}

	if pvcReady && stsReady && svcReady && ingressReady {
		addCondition(&devspace.Status, "Ready", metav1.ConditionTrue, "AllResourcesReady", "All resources are ready")
	} else {
		addCondition(&devspace.Status, "NotReady", metav1.ConditionFalse, "SomeResourcesNotReady", "Some resources are not ready")
	}

	log.Info("Reconciliation Complete...")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *DevSpaceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	r.Recorder = mgr.GetEventRecorderFor("devspace-controller")

	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.DevSpace{}).
		Complete(r)
}
