# The output binary name
BINARY_NAME=judoka
OUTPUT_DIR=bin
# Go tools
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

all: build

build: 
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v

build-wasm:	
	GOOS=js GOARCH=wasm $(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME).wasm
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js $(OUTPUT_DIR)
	cp provision/wasm/index.html $(OUTPUT_DIR)

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -fr $(OUTPUT_DIR)

run:
	$(GORUN) .