VERSION := 0.1.0

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/clever -ldflags "-X main.version=$(VERSION)"

.PHONY: build-linux
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/clever-linux -ldflags "-X main.version=$(VERSION)"

.PHONY: test
test:
	go test -v

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)

