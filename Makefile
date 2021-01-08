.PHONY: build
build:  
	go build -o build -v ./cmd/apiserver
	
.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL := build

