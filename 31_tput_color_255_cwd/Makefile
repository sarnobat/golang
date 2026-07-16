BINARY := cwd2color
ARCH := $(shell go env GOARCH)
UNAME := $(shell uname)

.PHONY: all build build-linux build-mac clean run

all: build

build: build-linux build-mac

build-linux:
	GOOS=linux GOARCH=$(ARCH) go build -o $(BINARY).Linux

build-mac:
	GOOS=darwin GOARCH=$(ARCH) go build -o $(BINARY).Darwin

run: build
	./$(BINARY).$(UNAME)

clean:
	rm -f $(BINARY) $(BINARY).Linux $(BINARY).Darwin
