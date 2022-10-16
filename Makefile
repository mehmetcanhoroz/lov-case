install:
	cp .env.local .env
	go mod download
build-client:
	make install
	./scripts/build.sh client client/main.go client
build-server:
	make install
	./scripts/build.sh server main.go server
build:
	make install
	./scripts/build.sh $(fn) $(from) $(where)
lint:
	golangci-lint run ./... --fast

clean:
	rm -rf ./bin Gopkg.lock