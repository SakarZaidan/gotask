BINARY_NAME=gotask

build:
	go build -o $(BINARY_NAME) ./cmd/gotask/main.go

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

install:
	go install ./cmd/gotask

.PHONY: build test clean install
