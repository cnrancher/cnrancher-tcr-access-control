#!/usr/bin/env bash

set -euo pipefail

cd $(dirname $0)/../
WORKINGDIR=$(pwd)

cd $WORKINGDIR/scripts
./validate.sh
./test.sh
./build.sh
