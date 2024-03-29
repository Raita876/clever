VERSION := 0.1.1

.PHONY: build
build:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/clever -ldflags "-X main.version=$(VERSION)"
	GOOS=linux GOARCH=amd64 go build -o ./bin/clever-linux -ldflags "-X main.version=$(VERSION)"

.PHONY: test
test:
	go test -v

.PHONY: tag
tag:
	git tag $(VERSION)
	git push origin $(VERSION)

