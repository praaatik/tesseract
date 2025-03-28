.PHONY: build test bench clean cover

build:
	go build -o bin/tesseract ./main.go

run:
	 CUBE_HOST=localhost CUBE_PORT=5555 ./bin/tesseract

run-debug:
	CUBE_HOST=localhost CUBE_PORT=5555 ./bin/tesseract -loglevel=DEBUG

run-error:
	CUBE_HOST=localhost CUBE_PORT=5555 ./bin/tesseract -loglevel=ERROR

run-warn:
	CUBE_HOST=localhost CUBE_PORT=5555 ./bin/tesseract -loglevel=WARN

test:
	go test ./... -v

cover:
	go test ./... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

bench:
	go test -bench=. ./...

clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

help:
	@echo "Available commands:"
	@echo "  make build             - Build all binaries."
	@echo "  make test              - Run all tests."
	@echo "  make cover             - Run tests with coverage."
	@echo "  make bench             - Run all benchmarks."
	@echo "  make clean             - Clean build artifacts."
	@echo "  make run               - Run the binary (defaults to info logging mode)."
	@echo "  make run-debug         - Run the binary in debug logging mode."
	@echo "  make run-error         - Run the binary in error logging mode."
	@echo "  make run-warn          - Run the binary in warn logging mode."
