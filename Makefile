default: build

build:
	@go fmt
	@go vet
	@golint
	@go test ./...
	@go build

cover:
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out
