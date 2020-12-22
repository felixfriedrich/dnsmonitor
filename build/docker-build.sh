#!/usr/bin/env bash
set -e
# set -x

source "$(dirname "${BASH_SOURCE[0]}")/.common.sh"

env CGO_ENABLED=0 go build -a -installsuffix cgo -o ./build/dnsmonitor ./cmd/dnsmonitor/main.go

DOCKER_BUILD_COMMAND="docker build"
DOCKER_BUILD_COMMAND="${DOCKER_BUILD_COMMAND} -t ${NAME}:latest"
DOCKER_BUILD_COMMAND="${DOCKER_BUILD_COMMAND} -t ${NAME}:${VERSION}"

# Tags for registries
for REGISTRY in "${REGISTRIES[@]}"; do
    DOCKER_BUILD_COMMAND="${DOCKER_BUILD_COMMAND} -t ${REGISTRY}${NAME}:latest"
    DOCKER_BUILD_COMMAND="${DOCKER_BUILD_COMMAND} -t ${REGISTRY}${NAME}:${VERSION}"
done

DOCKER_BUILD_COMMAND="$DOCKER_BUILD_COMMAND ./build"

eval "${DOCKER_BUILD_COMMAND}"

echo
echo "docker images"
echo "-------------"
docker images | grep dnsmonitor
