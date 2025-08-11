current_dir = $(shell pwd)
test:
	go test -v -cover -short ./...

.PHONY: test