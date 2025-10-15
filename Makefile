export TAG := $(shell git describe --tags --always)

.PHONY: build
build:
	go build -o build/app cmd/app/*.go

.phony: docker-build
docker-build: build
	IMAGE_TAG=$(TAG) docker buildx bake

.phony: image
image: docker-build
	mkdir -p image/
	docker save seddy.com/go-getting-started:$(TAG) -o image/go-getting-started-$(TAG).tar

.PHONY: clean
clean:
	rm -rf build/
	rm -rf image/

