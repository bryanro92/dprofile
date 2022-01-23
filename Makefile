SHELL = /bin/bash
COMMIT = $(shell git rev-parse --short HEAD)

install: build 
	chmod 555 ./dprofile

build:
	go build -ldflags '-X github.com/bryanro92/dprofile/pkg/version.GitCommit=$(COMMIT)' . 

run-clone:
	go run -ldflags '-X github.com/bryanro92/dprofile/pkg/version.GitCommit=$(COMMIT)' . clone

clean:  
	rm -f ./dprofile

.PHONY: build install run-clone clean

