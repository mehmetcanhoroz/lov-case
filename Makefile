build-client:
	go mod download
	env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/client ./client/main.go

lint:
	golangci-lint run ./... --fast
