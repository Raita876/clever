.PHONY: build
build:
	go build -o clever

.PHONY: test
test:
	go test -v