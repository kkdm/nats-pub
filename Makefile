VERSION=v0.1.0

build:
	go build -o ./bin  -ldflags="-s -w -X main.version=$(VERSION)" -trimpath ./src/nats-pub

test:
	go test ./src/nats-pub
