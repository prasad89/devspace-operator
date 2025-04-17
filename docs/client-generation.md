

## Client Generation

```bash 
❯ export CODEGEN_PKG=/Users/vishal/worktest/code-generator
❯ bash "${CODEGEN_PKG}"/generate-groups.sh "all" pkg/client github.com/prasad89/devspace-operator/pkg/apis "platform.devspace.io:v1alpha1" --go-header-file "$(pwd)/hack/boilerplate.go.txt"

Generating deepcopy funcs
Generating clientset for platform.devspace.io:v1alpha1 at pkg/client/clientset
Generating listers for platform.devspace.io:v1alpha1 at pkg/client/listers
Generating informers for platform.devspace.io:v1alpha1 at pkg/client/informers
```