install:
	cp .env.local .env
	go mod download
build-client:
	make build fn=client from=client/main.go where=client
build-server:
	make build fn=server from=main.go where=server
build:
	make install
	./scripts/build.sh $(fn) $(from) $(where)
lint:
	golangci-lint run ./... --fast

clean:
	rm -rf ./bin Gopkg.lock