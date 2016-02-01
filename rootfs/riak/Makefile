# Short name: Short name, following [a-zA-Z_], used all over the place.
# Some uses for short name:
# - Docker image name
# - Kubernetes service, rc, pod, secret, volume names
SHORT_NAME := riak

# SemVer with build information is defined in the SemVer 2 spec, but Docker
# doesn't allow +, so we use -.
VERSION ?= git-$(shell git rev-parse --short HEAD)

# Legacy support for DEV_REGISTRY, plus new support for DEIS_REGISTRY.
DEV_REGISTRY ?= $(eval docker-machine ip deis):5000
DEIS_REGISTY ?= ${DEV_REGISTRY}/
IMAGE_PREFIX ?= deis

IMAGE := ${DEIS_REGISTRY}${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}

all: build docker-build docker-push

# For cases where we're building from local
# We also alter the RC file to set the image name.
docker-build:
	docker build --rm -t ${IMAGE} ${CURDIR}

# Push to a registry that Kubernetes can access.
docker-push:
	docker push ${IMAGE}

.PHONY: all build docker-compile kube-up kube-down deploy
