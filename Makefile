NAME = "glusterfs-csi-driver"

.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/$(NAME) ./cmd/glusterfs
