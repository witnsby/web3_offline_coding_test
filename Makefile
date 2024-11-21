SHELL=bash
GOOS?=darwin
GOARCH?=arm64

.PHONY: tests
tests:
	@set -e

	go test ./... -v || exit 1; \

.PHONY: cover
cover:
	@set -e

	go test -short -count=1 -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm ./coverage.out

.PHONY: build-app
build-app:
	@set -e

	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./tmp/rock-paper-scissors ./src/cmd/bin/main.go || exit 1; \
