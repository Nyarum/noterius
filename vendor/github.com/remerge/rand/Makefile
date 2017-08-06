PROJECT := go-rand
PACKAGE := github.com/remerge/$(PROJECT)

# http://stackoverflow.com/questions/322936/common-gnu-makefile-directory-path#comment11704496_324782
TOP := $(dir $(CURDIR)/$(word $(words $(MAKEFILE_LIST)),$(MAKEFILE_LIST)))

GOFMT=gofmt -w -s

GOSRCDIR=$(GOPATH)/src/$(PACKAGE)
GOPATHS=$(shell glide novendor)
GOFILES=$(shell git ls-files | grep '\.go$$')

.PHONY: build clean lint test bench fmt dep init up gen

all: build

build: fmt
	cd $(GOSRCDIR) && \
		CGO_ENABLED=0 \
		go build -v

clean:
	go clean
	rm -rf $(TOP)/vendor/

lint:
	cd $(GOSRCDIR) && \
		gometalinter --deadline=60s --enable-all --errors --tests $(GOPATHS)

test: lint
	cd $(GOSRCDIR) && \
		go test -timeout 60s -v $(GOPATHS)

bench:
	cd $(GOSRCDIR) && \
		go test -bench=. -cpu 4 $(GOPATHS)

fmt:
	$(GOFMT) $(GOFILES)

dep:
	go get -u github.com/Masterminds/glide
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install --update
	cd $(GOSRCDIR) && glide install

init:
	cd $(GOSRCDIR) && \
		glide init

up:
	cd $(GOSRCDIR) && \
		glide update

gen:
	cd $(GOSRCDIR) && \
		go generate $(GOPATHS)
	$(GOFMT) $(GOFILES)
