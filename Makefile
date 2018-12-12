.PHONY: lint vet test build
lint:
	@golint $(shell go list ./...|grep -v vendor)

vet:
	@go vet ./...

test:
	@go test -v $(shell go list ./... | grep -v /vendor/)
	
build:
	@go build -o gocron .