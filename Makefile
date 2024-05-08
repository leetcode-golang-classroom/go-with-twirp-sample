.PHONY=build

gen-client:
	@protoc --twirp_out=paths=source_relative:. --go_out=paths=source_relative:. rpc/haberdasher/service.proto

build-server:
	@go build -o bin/server cmd/server/main.go

run-server: build-server
	@./bin/server

build-client:
	@go build -o bin/client cmd/client/main.go

run-client: build-client
	@./bin/client

test:
	@go test -v -cover ./test/...