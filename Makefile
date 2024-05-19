build:
	go build -o bin/GoAPI ./cmd
	
run: build
	@./bin/GoAPI

test:
	@go test -v  ./cmd/...
