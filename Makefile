default: build

lint:
	go fmt ./cmd
	go vet ./cmd
	gometalinter --deadline=15s ./cmd/...

build:
	go fmt ./cmd
	go vet ./cmd
	go build -o dsql ./cmd/
test: build
	go test -v ./cmd/...
   zip: build	
	zip dsql-linux-$(shell arch).zip dsql