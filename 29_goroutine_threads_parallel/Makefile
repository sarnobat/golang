# Name of the binary
BINARY_NAME=main

# Output directory
BUILD_DIR=build

# List of target OS/ARCH pairs
TARGETS = \
	darwin/arm64 \
	darwin/amd64 \
	linux/arm64 \
	linux/amd64 \
	windows/amd64 \
	windows/arm64

# Default target: build all
all: $(TARGETS)

# Rule to build for each OS/ARCH
$(TARGETS):
	@osarch=$@; \
	GOOS=$${osarch%/*}; \
	GOARCH=$${osarch#*/}; \
	ext=""; \
	if [ "$${GOOS}" = "windows" ]; then ext=".exe"; fi; \
	output="$(BUILD_DIR)/$${GOOS}_$${GOARCH}/$(BINARY_NAME)$${ext}"; \
	echo "→ Building for $${GOOS}/$${GOARCH} → $${output}"; \
	mkdir -p "$$(dirname $$output)"; \
	GOOS=$${GOOS} GOARCH=$${GOARCH} go build -o "$$output" .

# Clean build artifacts
clean:
	rm -rf $(BUILD_DIR)

.PHONY: all clean $(TARGETS)
