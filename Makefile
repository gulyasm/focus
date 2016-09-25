default: build

installdep:
	go get -u github.com/golang/lint/golint
	go get -u

build: installdep
	@go fmt
	@go vet
	@golint
	@go test ./...
	@go build

cover: installdep
	@go tool cover -func=coverage.out
	@go tool cover -html=coverage.out
	@rm coverage.out
