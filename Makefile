build:
	go build -o bin/libpostal-grpc -ldflags "-s -w" cmd/libpostal-grpc/main.go

run:
	go run cmd/libpostal-grpc/main.go

test:
	go test ./... -v