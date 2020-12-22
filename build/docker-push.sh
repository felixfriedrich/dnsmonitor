#!/usr/bin/env bash
set -e
# set -x

source "$(dirname "${BASH_SOURCE[0]}")/.common.sh"

for REGISTRY in "${REGISTRIES[@]}"; do
    PUSH_COMMAND="docker push ${REGISTRY}${NAME}:latest"
    eval "${PUSH_COMMAND}"
done
