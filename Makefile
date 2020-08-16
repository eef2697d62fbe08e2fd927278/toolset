GOPATH = $(shell pwd)

build_website:
	mkdir -p $(GOPATH)/bin
	go build -o $(GOPATH)/bin/website -a $(GOPATH)/cmd/website

build_api:
	mkdir -p $(GOPATH)/bin
	go build -o $(GOPATH)/bin/api -a $(GOPATH)/cmd/api

build_all:
	$(MAKE) build_frontend
	$(MAKE) build_backend
