GOPATH = .

build_frontend:
	go build -a $(GOPATH)/cmd/frontend -o $(GOPATH)/bin/frontend.exe

build_backend:
	go build -a $(GOPATH)/cmd/backend -o $(GOPATH)/bin/backend.exe

build_all:
	$(MAKE) build_frontend
	$(MAKE) build_backend
