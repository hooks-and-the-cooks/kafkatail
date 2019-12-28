NOVENDOR=$(shell go list ./... | grep -v /vendor/)

setup:
	go get -u golang.org/x/lint/golint

build:
	GO111MODULE=on go mod vendor

compile:
	GO111MODULE=on go build $(NOVENDOR)

test:
	GO111MODULE=on go test $(NOVENDOR)
