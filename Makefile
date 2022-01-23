SHELL = /bin/bash
COMMIT = $(shell git rev-parse --short=7 HEAD)

install: build 
	chmod 555 ./dprofile
	# mv ./dprofile /usr/local/bin/	

build:
	go build . -ldflags "-X github.com/bryanro92/dprofile/pkg/util/version.GitCommit=$(COMMIT)"

.PHONY: build install
