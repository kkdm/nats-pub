VERSION=v0.1.0

build:
	go build -o ./bin/nats-pub  -ldflags="-s -w -X main.version=$(VERSION)" -trimpath ./cmd/nats-pub

test:
	go test ./cmd/nats-pub
