################################################################################
#### INSTALLATION VARS
################################################################################
PREFIX=/usr/local

################################################################################
#### BUILD VARS
################################################################################
BIN=seqf
BINBETA=renum-beta
BINS=.
HEAD=$(shell git describe --dirty --long --tags 2> /dev/null  || git rev-parse --short HEAD)
TIMESTAMP=$(shell date '+%Y-%m-%dT%H:%M:%S %z %Z')

LDFLAGS="-X 'main.buildVersion=$(HEAD)' -X 'main.buildTimestamp=$(TIMESTAMP)' -X 'main.compiledBy=$(shell go version)'" # `-s -w` removes some debugging info that might not be necessary in production (smaller binaries)

all: local

################################################################################
#### HOUSE CLEANING
################################################################################

clean:
	rm -f $(BIN) $(BIN)-* $(BINS)/$(BIN) $(BINS)/$(BIN)-*

.PHONY: mod
mod:
	go mod tidy
	go mod vendor

.PHONY: check
check: mod
	golint
	goimports -w ./
	gofmt -w ./
	go vet

################################################################################
#### INSTALL
################################################################################

.PHONY: install
install:
	mkdir -p $(PREFIX)/bin
	cp $(BINS)/$(BIN) $(PREFIX)/bin/$(BIN)

.PHONY: install-beta
install-beta:
	mkdir -p $(PREFIX)/bin
	cp $(BINS)/$(BINBETA) $(PREFIX)/bin/$(BINBETA)

.PHONY: uninstall
uninstall:
	rm -f $(PREFIX)/bin/$(BIN) $(PREFIX)/bin/$(BINBETA)

################################################################################
#### ENV BUILDS
################################################################################

.PHONY: local
local: check
	go build -ldflags $(LDFLAGS) -o $(BINS)/$(BIN)

.PHONY: beta
beta: check
	go build -ldflags $(LDFLAGS) -o $(BINS)/$(BINBETA)

.PHONY: local-vendor
local-vendor: check
	go build -mod=vendor -ldflags $(LDFLAGS) -o $(BINS)/$(BIN)

.PHONY: test
test: mod
	go test -mod=vendor -coverprofile=coverage.out -covermode=count ./...
