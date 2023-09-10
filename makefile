# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get

# Name of your binary executable
BINARY_NAME := practice

# Main build target
all: build

# Build the binary
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Clean the project
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Install Go dependencies
get:
	$(GOGET)

# Run tests
test:
	$(GOTEST) ./...

# Run the binary
run: 
	./$(BINARY_NAME)

build-run: build run

.PHONY: all build clean get test run
