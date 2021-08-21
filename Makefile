export version=$(shell cat VERSION)
export LDFLAGS=-ldflags "-X main.BuilderVersion=$(version)"

.PHONY: all
all: build

clean:
	# Clean the earlier build
	rm -f bin/detect bin/detect bin/release

build: server-build

server-build: format-check clean
	# Build the source code
	go build -o bin/detect $(LDFLAGS) ./cmd/server/detect
	go build -o bin/build $(LDFLAGS) ./cmd/server/build
	go build -o bin/release $(LDFLAGS) ./cmd/server/release

server-build-alpine: format-check clean
	# Build the source code
	GOOS=linux GOARCH=amd64 go build -o bin/detect $(LDFLAGS) ./cmd/server/detect
	GOOS=linux GOARCH=amd64 go build -o bin/build $(LDFLAGS) ./cmd/server/build
	GOOS=linux GOARCH=amd64 go build -o bin/release $(LDFLAGS) ./cmd/server/release

test:
	CGO_ENABLED=0 go test -v ./... -p 1 -count=1

format:
	go fmt ./...

format-check: export format_output=$(shell gofmt -l .)
format-check:
	# Check format correctness
	@[ "${format_output}" == "" ] || exit -1
