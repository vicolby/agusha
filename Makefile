build:
	go build -o bin/dswv3

run:
	go run ./cmd/main.go

test:
	go test -v ./...
