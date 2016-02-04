default: build

lint:
	go fmt
	go vet
	gometalinter --deadline=15s ./...

build:
	go fmt
	go vet
	go build
test: build
	go test -v ./...
coverage-test:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out
