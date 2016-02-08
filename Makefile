# Short name: Short name, following [a-zA-Z_], used all over the place.
# Some uses for short name:
# - Docker image name
# - Kubernetes service, rc, pod, secret, volume names
SHORT_NAME := riak

# SemVer with build information is defined in the SemVer 2 spec, but Docker
# doesn't allow +, so we use -.
VERSION ?= git-$(shell git rev-parse --short HEAD)

# Docker Root FS
BINDIR := ./rootfs
BINARY_DEST_DIR := ${BINDIR}/bin
LDFLAGS := "-s -X main.version=${VERSION}"

# Dockerized development environment variables
REPO_PATH := github.com/deis/${SHORT_NAME}
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.3.0
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -e GO15VENDOREXPERIMENT=1 -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}

# Legacy support for DEV_REGISTRY, plus new support for DEIS_REGISTRY.
DEV_REGISTRY ?= $(eval docker-machine ip deis):5000
DEIS_REGISTY ?= ${DEV_REGISTRY}/
IMAGE_PREFIX ?= deis

# Kubernetes-specific information for RC, Service, and Image.
BOOTSTRAP := manifests/deis-${SHORT_NAME}-bootstrap-pod.yaml
RC := manifests/deis-${SHORT_NAME}-rc.yaml
SVC := manifests/deis-${SHORT_NAME}-service.yaml
DISCO_SVC := manifests/deis-${SHORT_NAME}-discovery-service.yaml
CLUSTER_SVC := manifests/deis-${SHORT_NAME}-cluster-service.yaml

IMAGE := ${DEIS_REGISTRY}${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}

TEST_PACKAGES := $(shell ${DEV_ENV_CMD} glide nv)

all: build docker-build docker-push

bootstrap:
		${DEV_ENV_CMD} glide install

glideup:
		${DEV_ENV_CMD} glide up

build:
	${DEV_ENV_PREFIX} -e CGO_ENABLED=0 ${DEV_ENV_IMAGE} go build -a -installsuffix cgo -ldflags ${LDFLAGS} -o ${BINARY_DEST_DIR}/boot boot.go

test:
	${DEV_ENV_CMD} go test -race ${TEST_PACKAGES}

docker-build:
	docker build --rm -t ${IMAGE} rootfs

# Push to a registry that Kubernetes can access.
docker-push:
	docker push ${IMAGE}

# Deploy is a Kubernetes-oriented target
deploy: kube-service kube-rc

# Some things, like services, have to be deployed before pods. This is an
# example target. Others could perhaps include kube-secret, kube-volume, etc.
kube-service:
	kubectl create -f ${SVC}
	kubectl create -f ${DISCO_SVC}
	kubectl create -f ${CLUSTER_SVC}

# When possible, we deploy with RCs.
kube-rc:
	kubectl create -f ${BOOTSTRAP}
	kubectl create -f ${RC}

# We don't need to delete the bootstrap pod here because it'll get selected by the RC.
kube-clean:
	kubectl delete -f ${RC}
	kubectl delete -f ${SVC}
	kubectl delete -f ${DISCO_SVC}
	kubectl delete -f ${CLUSTER_SVC}

.PHONY: all build docker-compile kube-up kube-down deploy
