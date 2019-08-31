VERSION := 0.1.0

.PHONY: build
build:
	go build -o clever -ldflags "-X main.version=$(VERSION)"

.PHONY: test
test:
	go test -v