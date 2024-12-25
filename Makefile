.PHONY: all
all: build

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build
build: vet
	go build -o bin/jot jot.go

.PHONY: run
run: vet
	go run main.go

.PHONY: install
install: vet
	go install .
