GOPATH = $(shell pwd)

build_website:
	mkdir -p $(GOPATH)/bin
	go build -a $(GOPATH)/cmd/website -o $(GOPATH)/bin

build_api:
	mkdir -p $(GOPATH)/bin
	go build $(GOPATH)/cmd/api

build_all:
	$(MAKE) build_frontend
	$(MAKE) build_backend
