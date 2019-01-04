# Go parameters
GOCMD=go
STATIKCMD=statik
NPMCMD=npm
NPMLINT=$(NPMCMD) run lint
NPMBUILD=$(NPMCMD) run build
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=gddu
BINARY_UNIX=$(BINARY_NAME)_unix
STATIC_DIR=dist
VENDOR_DIR=vendor/gddu
GORELEASER=goreleaser release --rm-dist

all: test build
build-deps: 
	$(NPMBUILD)
	$(STATIKCMD) -src=$(STATIC_DIR) -dest=$(VENDOR_DIR)
build: 
	$(NPMBUILD)
	$(STATIKCMD) -src=$(STATIC_DIR) -dest=$(VENDOR_DIR)
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(NPMLINT)
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

release:
	$(GORELEASER)


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v