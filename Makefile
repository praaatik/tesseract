.PHONY: build test bench clean cover

build:
	go build -o bin/tesseract ./main.go

run:
	./bin/tesseract

run-debug:
	./bin/tesseract -loglevel=DEBUG

run-error:
	./bin/tesseract -loglevel=ERROR

run-warn:
	./bin/tesseract -loglevel=WARN

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
	@echo "  make build         - Build all binaries"
	@echo "  make test          - Run all tests"
	@echo "  make cover         - Run tests with coverage"
	@echo "  make bench         - Run all benchmarks"
	@echo "  make clean         - Clean build artifacts"
	@echo "  make doc           - Generate documentation"
