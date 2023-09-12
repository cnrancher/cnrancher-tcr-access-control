#!/usr/bin/env bash

set -euo pipefail

cd $(dirname $0)/../
WORKINGDIR=$(pwd)

source ./scripts/version.sh
echo "Build version: ${VERSION}"

BUILD_FLAG="${BUILD_FLAG:-}"
if [[ -z "${DEBUG:-}" ]]; then
    BUILD_FLAG="-extldflags -static -s -w"
fi
if [[ "${COMMIT}" != "UNKNOW" ]]; then
    BUILD_FLAG="${BUILD_FLAG} -X 'github.com/cnrancher/tcr-access-control/pkg/utils.GitCommit=${COMMIT}'"
fi
BUILD_FLAG="${BUILD_FLAG} -X 'github.com/cnrancher/tcr-access-control/pkg/utils.Version=${VERSION}'"

mkdir -p bin && cd bin
CGO_ENABLED=0 go build -ldflags "${BUILD_FLAG}" -o tcr-access-control-${OS}-${ARCH} ..
echo "--------------------"
ls -alh tcr-access-control*
echo "--------------------"
