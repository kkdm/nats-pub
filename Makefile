VERSION=v0.1.0

build:
	go build -o ./bin/natspub  -ldflags="-s -w -X main.version=$(VERSION)" -trimpath ./src/natspub

test:
	go test ./src/natspub
