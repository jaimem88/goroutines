.PHONY: run build run-docker

run:
	go run ./cmd/goroutines/main.go

build:
	go build -o ./bin/goroutines ./cmd/goroutines/

run-docker:
	docker run --rm -v ${PWD}:/go/app -w /go/app golang:1.13.4 go run ./cmd/goroutines/main.go
