SHORT_NAME := riak

VERSION ?= git-$(shell git rev-parse --short HEAD)
DEV_REGISTRY ?= $(eval docker-machine ip deis):5000
DEIS_REGISTY ?= ${DEV_REGISTRY}/
IMAGE_PREFIX ?= deis
IMAGE := ${DEIS_REGISTRY}${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}

docker-build:
	docker build --rm -t ${IMAGE} ${CURDIR}

docker-push:
	docker push ${IMAGE}
