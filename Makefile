test:
	go test ./...

.PHONY: examples
examples:
	go run ./examples/01-basic-project/.caye
