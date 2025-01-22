#!/bin/env bash

DOCKER_IMAGE_NAME=golang:1.23-bullseye


docker run -w /build/ --name build_noverify -v $(pwd):/build:rw -d ${DOCKER_IMAGE_NAME} sleep 3000
docker exec build_noverify bash -c 'git config --global --add safe.directory /build'
docker exec build_noverify bash -c 'make build'
docker rm -vf build_noverify

