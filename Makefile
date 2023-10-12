
default: build

generate:
	go install github.com/golang/mock/mockgen@v1.6.0
	go generate ./...

build: generate
	go build

install: generate
	go install ./cmd/mailtick

test: generate
	go test ./...
