.PHONY: all build run test clean

BINARY_NAME=gotrova

all: test build

build:
	go build -o $(BINARY_NAME) ./cmd/gotrova

run: build
	./$(BINARY_NAME)

test:
	go test ./...

clean:
	go clean
	rm -f $(BINARY_NAME)
