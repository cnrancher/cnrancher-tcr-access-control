#!/usr/bin/env bash

set -euo pipefail

# Ensure git installed
type git > /dev/null

DIRTY=""
if [[ -n "$(git status --porcelain --untracked-files=no)" ]]; then
    DIRTY="-dirty"
fi

COMMIT=$(git rev-parse --short HEAD)
GIT_TAG=${DRONE_TAG:-$(git tag -l --contains HEAD | head -n 1)}

if [[ -z "$DIRTY" && -n "$GIT_TAG" ]]; then
    VERSION=$GIT_TAG
else
    VERSION="v0.0.0${DIRTY}"
fi

if [[ -z "${ARCH:-}" ]]; then
    ARCH=$(go env GOHOSTARCH)
fi
if [[ -z "${OS:-}" ]]; then
    OS=$(go env GOHOSTOS)
fi
