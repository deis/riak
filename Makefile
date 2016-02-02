SHORT_NAME := riak

VERSION ?= git-$(shell git rev-parse --short HEAD)

# build information
BINARY_DEST_DIR := rootfs/bin
LDFLAGS := "-s -X main.version=${VERSION}"
TEST_PACKAGES := $(shell ${DEV_ENV_CMD} glide nv)

# Dockerized development environment variables
REPO_PATH := github.com/deis/${SHORT_NAME}
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.3.0
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -e GO15VENDOREXPERIMENT=1 -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}

# Kubernetes resources
MANIFESTS_DIR := ${CURDIR}/manifests
BOOTSTRAP := ${MANIFESTS_DIR}/deis-riak-bootstrap-pod.yaml
RC := ${MANIFESTS_DIR}/deis-riak-rc.yaml
CS_RC := ${MANIFESTS_DIR}/deis-riak-cs-rc.yaml
STANCHION_RC := ${MANIFESTS_DIR}/deis-riak-stanchion-rc.yaml
SVC := ${MANIFESTS_DIR}/deis-riak-service.yaml
DISCO_SVC := ${MANIFESTS_DIR}/deis-riak-discovery-service.yaml
CLUSTER_SVC := ${MANIFESTS_DIR}/deis-riak-cluster-service.yaml
CS_SVC := ${MANIFESTS_DIR}/deis-riak-cs-service.yaml
STANCHION_SVC := ${MANIFESTS_DIR}/deis-riak-stanchion-service.yaml
ADMIN_USER_SEC := ${MANIFESTS_DIR}/deis-riak-cs-admin-user-secret.yaml

all: build riak-docker-build riak-cs-docker-build riak-stanchion-docker-build riak-docker-push riak-cs-docker-push riak-stanchion-docker-push

bootstrap:
		${DEV_ENV_CMD} glide install

glideup:
		${DEV_ENV_CMD} glide up

build:
	${DEV_ENV_PREFIX} -e CGO_ENABLED=0 ${DEV_ENV_IMAGE} go build -a -installsuffix cgo -ldflags ${LDFLAGS} -o ${BINARY_DEST_DIR}/boot boot.go
	# copy the built binary into each Docker filesystem, so the each Dockerfile can ADD it into the respective images
	mkdir -p rootfs/riak/bin
	mkdir -p rootfs/riak-cs/bin
	mkdir -p rootfs/riak-stanchion/bin
	cp ${BINARY_DEST_DIR}/boot rootfs/riak/bin/boot
	cp ${BINARY_DEST_DIR}/boot rootfs/riak-cs/bin/boot
	cp ${BINARY_DEST_DIR}/boot rootfs/riak-stanchion/bin/boot

test:
	${DEV_ENV_CMD} go test -race ${TEST_PACKAGES}

riak-build:
	make -C rootfs/riak build

riak-docker-build:
	make -C rootfs/riak docker-build

riak-docker-push:
	make -C rootfs/riak docker-push

riak-cs-docker-build:
	make -C rootfs/riak-cs docker-build

riak-cs-docker-push:
	make -C rootfs/riak-cs docker-push

riak-stanchion-docker-build:
	make -C rootfs/riak-stanchion docker-build

riak-stanchion-docker-push:
	make -C rootfs/riak-stanchion docker-push

# Deploy is a Kubernetes-oriented target
kube-deploy: kube-secrets kube-service kube-rc

kube-secrets:
	kubectl create -f ${ADMIN_USER_SEC}

kube-service:
	kubectl create -f ${SVC}
	kubectl create -f ${DISCO_SVC}
	kubectl create -f ${CLUSTER_SVC}
	kubectl create -f ${CS_SVC}
	kubectl create -f ${STANCHION_SVC}

kube-rc:
	kubectl create -f ${BOOTSTRAP}
	kubectl create -f ${RC}
	kubectl create -f ${STANCHION_RC}
	kubectl create -f ${CS_RC}

kube-clean:
	kubectl delete -f ${RC}
	kubectl delete -f ${STANCHION_RC}
	kubectl delete -f ${CS_RC}
	kubectl delete -f ${SVC}
	kubectl delete -f ${DISCO_SVC}
	kubectl delete -f ${CLUSTER_SVC}
	kubectl delete -f ${CS_SVC}
	kubectl delete -f ${STANCHION_SVC}
	kubectl delete -f ${ADMIN_USER_SEC}
