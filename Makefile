export TAG := $(shell git describe --tags --always)

.PHONY: build
build:
	go build -o build/app cmd/app/*.go

.phony: docker-build
docker-build: build
	IMAGE_TAG=$(TAG) docker buildx bake
