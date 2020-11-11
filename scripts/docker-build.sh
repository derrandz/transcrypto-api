#!/usr/bin/env bash

docker build \
    -f build/docker/Dockerfile . \
    -t summitto/txsigner:${DOCKER_TAG} \
    --build-arg PORT=$PORT \
    --build-arg DAEMON_PRIVATE_KEY=${DAEMON_PRIVATE_KEY} \
    --build-arg SERVICE_NAME=${SERVICE_NAME}
