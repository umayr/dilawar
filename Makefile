.PHONY: default deps clean fmt pretest lint lint-test list vet build test all
SHELL := /bin/bash
BINARY=dilawar

VERSION=0.1.0
BUILD_TIME=`date +%FT%T%z`

BRANCH=`git rev-parse --abbrev-ref HEAD`
COMMIT=`git rev-parse --short HEAD`

LDFLAGS="-X ${BINARY}.version=${VERSION} -X ${BINARY}.buildtime=${BUILD_TIME} -X ${BINARY}.branch=${BRANCH} -X ${BINARY}.commit=${COMMIT}"

GLIDE := $(shell glide --version 2>/dev/null)
GOX := $(shell which gox 2>/dev/null)

OSES := "linux windows darwin"
ARCHS := "amd64 386"

default: build

deps:
ifdef GLIDE
	@glide install
else
	@echo "Glide is not installed."
	@echo 'Reflex is not installed'
	@echo 'Installing Glide...'
	@go get github.com/Masterminds/glide
	@glide install
endif

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

pretest:
	@gofmt -d $$(find . -type f -name '*.go' -not -path "./vendor/*") 2>&1 | read; [ $$? == 1 ]

lint-test:
	@go get -v github.com/golang/lint/golint
	@golint ./... | grep -v vendor/ 2>&1

vet:
	@go vet $(go list -f '{{ .ImportPath }}' ./... | grep -v vendor/)

test: pretest vet lint-test
	@go test -v $$(go list -f '{{ .ImportPath }}' ./... | grep -v vendor/) -p=1

build: clean deps test
	@go build -x -ldflags ${LDFLAGS} -o bin/${BINARY} github.com/umayr/${BINARY}/cmd/${BINARY}

fmt:
	@gofmt -w $$(find . -type f -name '*.go' -not -path "./vendor/*")

lint:
	@go get -v github.com/golang/lint/golint
	@golint ./... | grep -v vendor/

list:
	@go list -f '{{ .ImportPath }}' ./... | grep -v vendor/

cross:
ifdef GOX
	@gox -ldflags=${LDFLAGS} -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}" -os=${OSES} -arch=${ARCHS} github.com/umayr/${BINARY}/cmd/${BINARY}
else
	@echo 'Gox is not installed'
	@echo 'Installing Gox...'
	@go get github.com/mitchellh/gox
	@gox -ldflags=${LDFLAGS} -output="bin/{{.Dir}}_{{.OS}}_{{.Arch}}" -os=${OSES} -arch=${ARCHS} github.com/umayr/${BINARY}/cmd/${BINARY}
endif