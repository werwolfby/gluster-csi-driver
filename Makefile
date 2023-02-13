NAME = "werwolfby/glusterfs-csi-driver"
VERSION=$(shell git describe --tags --always --dirty)
DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/$(NAME) ./cmd/glusterfs

.PHONY: container
container:
	@echo "Building container..."
	@docker build -t $(NAME):$(VERSION) --build-arg version=$(VERSION) --build-arg builddate=$(DATE) .
