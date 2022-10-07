.PHONY: build

PROJECT=power4
VERSION=$$(git rev-parse --short=10 HEAD)

clean:
	go clean -cache

proto:
	protoc --go_out=./pkg/infrastructure/proto --go-grpc_out=./pkg/infrastructure/proto pkg/infrastructure/proto/*.proto

build_client:
	go build -v cmd/clientMain.go

build_server:
	go build -v cmd/serverMain.go

build_all:
	go build -v cmd/clientMain.go
	go build -v cmd/serverMain.go


run:
	go run cmd/serverMain.go

container:
	docker build -f build/Dockerfile . -t $(PROJECT):$(VERSION)

container-run: container
	docker run -it --rm $(PROJECT):$(VERSION)