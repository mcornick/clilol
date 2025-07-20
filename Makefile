.PHONY: build
build:
	go build

.PHONY: clean
clean:
	go clean

.PHONY: fmt
fmt:
	gofumpt -w .

.PHONY: tags
tags:
	ctags *.go cmd/*.go

.PHONY: test
test:
	go test ./...

.PHONY: vet
vet:
	go vet ./...
