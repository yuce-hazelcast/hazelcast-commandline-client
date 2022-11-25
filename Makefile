.PHONY: build test test-cover view-cover

GIT_COMMIT = $(shell git rev-parse HEAD 2> /dev/null || echo unknown)
CLC_VERSION = $(shell git describe --tags `git rev-list --tags --max-count=1` || echo UNKNOWN)
LDFLAGS = "-X 'github.com/hazelcast/hazelcast-go-client/internal.ClientType=CLC' -X 'github.com/hazelcast/hazelcast-commandline-client/internal.GitCommit=$(GIT_COMMIT)' -X 'github.com/hazelcast/hazelcast-commandline-client/internal.Version=$(CLC_VERSION)' -X 'github.com/hazelcast/hazelcast-go-client/internal.ClientVersion=$(CLC_VERSION)'"
TEST_FLAGS ?= -v -count 1
COVERAGE_OUT = coverage.out
PACKAGES = $(shell go list ./... | grep -v go-prompt | grep -v termdbms | grep -v internal/it | tr '\n' ',')

build:
	go build -tags base,hazelcastinternal,hazelcastinternaltest -ldflags $(LDFLAGS)  -o build/clc ./cmd/clc

test:
	go test -tags base,hazelcastinternal,hazelcastinternaltest $(TEST_FLAGS) ./...

test-cover:
	go test -tags base,hazelcastinternal,hazelcastinternaltest $(TEST_FLAGS) -coverprofile=coverage.out -coverpkg $(PACKAGES) -coverprofile=$(COVERAGE_OUT) ./...

view-cover:
	go tool cover -func $(COVERAGE_OUT) | grep total:
	go tool cover -html $(COVERAGE_OUT) -o coverage.html
