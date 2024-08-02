.PHONY: all build run test clean

all: build

build:
	go build -o bin/main ./cmd/go_manage_my_files

run:
	go run ./cmd/go_manage_my_files

test:
	go test ./...

clean:
	rm -f bin/*
