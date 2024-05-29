build:# Define the name of the binary
BINARY_NAME=./bin/goapi

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

# Turn on Redis
db:
	docker compose up -d

# Run the binary
run: build
	$(BINARY_NAME)

# Clean up the binary
clean:
	rm -f $(BINARY_NAME)

