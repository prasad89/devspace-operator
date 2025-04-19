#!/usr/bin/env bash

set -euo pipefail

# Get the current Go module path, e.g., github.com/username/repo
THIS_PKG=$(go list -m)

PKG_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$PKG_ROOT"

# Split module path into components: host, user/org, and repo
# For example: github.com/prasad89/project -> github.com, prasad89, project
IFS='/' read -r MODULE_HOST MODULE_USER MODULE_REPO <<<"$THIS_PKG"

# Source the kube_codegen.sh script which contains codegen functions
source "${PKG_ROOT}/kube_codegen.sh"

# Create symlinks so the code-generator can resolve import paths correctly
[[ -L $MODULE_HOST ]] || ln -s . "$MODULE_HOST"
[[ -L $MODULE_USER ]] || ln -s .. "$MODULE_USER"

# Clean up symlinks on script exit
trap 'rm -f "$MODULE_HOST" "$MODULE_USER"' EXIT

# Run helper code generation
kube::codegen::gen_helpers "${THIS_PKG}"

# Run clientset, lister, and informer generation
kube::codegen::gen_client \
    --output-pkg "${THIS_PKG}/pkg/generated" \
    --output-dir "${THIS_PKG}/pkg/generated" \
    --with-watch \
    --with-applyconfig \
    "${THIS_PKG}"
