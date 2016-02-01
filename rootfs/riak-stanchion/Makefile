SHORT_NAME := riak-stanchion

DEV_REGISTRY ?= $(eval docker-machine ip deis):5000
DEIS_REGISTY ?= ${DEV_REGISTRY}/
IMAGE_PREFIX ?= deis
VERSION ?= git-$(shell git rev-parse --short HEAD)
IMAGE := ${DEIS_REGISTRY}${IMAGE_PREFIX}/${SHORT_NAME}:${VERSION}

docker-build:
	docker build --rm -t ${IMAGE} ${CURDIR}

docker-push:
	docker push ${IMAGE}
