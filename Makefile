tools:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/HEAD/install.sh | sh -s -- -b ./bin v2.0.2
	curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v30.2/protoc-30.2-linux-x86_64.zip && unzip -oq protoc-30.2-linux-x86_64.zip -d ./bin/protoc && rm -f protoc-30.2-linux-x86_64.zip

lint:
	./bin/golangci-lint run

test:
	go test -v -race ./...

