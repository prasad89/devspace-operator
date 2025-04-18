#!/usr/bin/env bash

set -euo pipefail

THIS_PKG="github.com/prasad89/devspace-operator"
PKG_ROOT=$(realpath "$(dirname "${BASH_SOURCE[0]}")/..")

# TODO: Make CODEGEN_PKG path configurable or auto-detectable
CODEGEN_PKG="/home/prasad89/code-generator"

cd "$PKG_ROOT"

source "${CODEGEN_PKG}/kube_codegen.sh"

# TODO: Replace symlinks with a more generic or configurable approach
ln -s . github.com
ln -s .. prasad89
trap 'rm -f github.com prasad89' EXIT

kube::codegen::gen_helpers \
    "${THIS_PKG}" \
    --boilerplate /dev/null

kube::codegen::gen_client \
    --output-pkg "${THIS_PKG}/pkg/generated" \
    --output-dir "${THIS_PKG}/pkg/generated" \
    --boilerplate /dev/null \
    --with-watch \
    --with-applyconfig \
    "${THIS_PKG}"
