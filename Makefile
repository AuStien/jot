all: build

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: vet fmt
	go build -o bin/log cmd/log/main.go

.PHONY: run
run: vet fmt
	go run cmd/log/main.go

