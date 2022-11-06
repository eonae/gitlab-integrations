.PHONY: build

build:
	go build -o build/gi src/*.go

start:
	go run src/*.go