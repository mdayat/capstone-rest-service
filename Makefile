all:
		test vet fmt build

run:
		go run cmd/main.go

test:
		go test ./...

vet:
		go vet ./...

fmt:
		go fmt ./...

build:
		go build -o bin/main ./cmd/main.go