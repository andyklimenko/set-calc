all: build test

test:
	go test -count=1 -race ./...
build:
	go build -o set-calc main.go