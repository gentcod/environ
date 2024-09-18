current_dir = $(shell pwd)
run:
	go run .

test:
	go test -v -cover -short ./...

.PHONY: run test