tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b ./bin v2.0.2

lint:
	./bin/golangci-lint run

test:
	go test -v -race ./...

protoc:
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v31.0/protoc-31.0-linux-x86_64.zip && unzip -oq protoc-31.0-linux-x86_64.zip -d ./bin/protoc && rm -f protoc-31.0-linux-x86_64.zip

protoc_gen:
	./bin/protoc/bin/protoc --go_out=. --go-grpc_out=. api/ping.proto

run:
	go run cmd/app/main.go