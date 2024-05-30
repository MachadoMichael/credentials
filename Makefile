build:# Define the name of the binary
BINARY_NAME=./bin/credentials

# Default target
all: build

# Build the binary
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

# Turn on Redis
db:
	@echo "Openning Docker Desktop..."
	open -a Docker &
	sleep 10
	@echo "Docker Desktop opened. Running docker compose"
	docker compose up -d

# Run the binary
run: build
	$(BINARY_NAME)

# Clean up the binary
clean:
	rm -f $(BINARY_NAME)

