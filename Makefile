# Short name: Short name, following [a-zA-Z_], used all over the place.
# Some uses for short name:
# - Docker image name
# - Kubernetes service, rc, pod, secret, volume names
SHORT_NAME := riak

# SemVer with build information is defined in the SemVer 2 spec, but Docker
# doesn't allow +, so we use -.
VERSION := 0.0.1-$(shell date "+%Y%m%d%H%M%S")

# Docker Root FS
BINDIR := ./rootfs

# Legacy support for DEV_REGISTRY, plus new support for DEIS_REGISTRY.
DEV_REGISTRY ?= $(eval docker-machine ip deis):5000
DEIS_REGISTY ?= ${DEV_REGISTRY}

# Kubernetes-specific information for RC, Service, and Image.
BOOTSTRAP := manifests/${SHORT_NAME}-bootstrap-pod.yaml
RC := manifests/${SHORT_NAME}-rc.yaml
SVC := manifests/${SHORT_NAME}-service.yaml
DISCO_SVC := manifests/${SHORT_NAME}-discovery-service.yaml
IMAGE := ${DEIS_REGISTRY}/deis/${SHORT_NAME}:${VERSION}

all: docker-build docker-push

# For cases where we're building from local
# We also alter the RC file to set the image name.
docker-build:
	docker build --rm -t ${IMAGE} rootfs
	perl -pi -e "s|[a-z0-9.:]+\/deis\/${SHORT_NAME}:[0-9a-z-.]+|${IMAGE}|g" ${BOOTSTRAP}
	perl -pi -e "s|[a-z0-9.:]+\/deis\/${SHORT_NAME}:[0-9a-z-.]+|${IMAGE}|g" ${RC}

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

# When possible, we deploy with RCs.
kube-rc:
	kubectl create -f ${BOOTSTRAP}
	kubectl create -f ${RC}

# We don't need to delete the bootstrap pod here because it'll get selected by the RC.
kube-clean:
	kubectl delete -f ${RC}
	kubectl delete -f ${SVC}
	kubectl delete -f ${DISCO_SVC}

.PHONY: all build docker-compile kube-up kube-down deploy
