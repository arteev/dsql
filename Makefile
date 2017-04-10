default: build

lint:
	go fmt 
	go vet 
	gometalinter --deadline=15s /...

build:
	go fmt 
	go vet 
	go build -o dsql 
test: build
	go test -v ./...
   zip: build	
	zip dsql-linux-$(shell arch).zip dsql