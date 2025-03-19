# Default target
.PHONY: all
all: build

# Build the gadget binary
.PHONY: build
build:
	go build -o gadget ./cmd/gadget

# Install the gadget binary to the system
.PHONY: install
install:
	go install ./cmd/gadget

# Clean up generated files
.PHONY: clean
clean:
	rm -f gadget