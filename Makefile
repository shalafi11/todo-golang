GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get


build: 
		$(GOBUILD) -o $(BINARY_NAME) -v
test: 
		$(GOTEST) -v ./... 
run:
		$(GOCMD) run main.go