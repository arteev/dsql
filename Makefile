VERSION="1.6.25"
BUILD_TIME=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
GITHEAD=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X github.com/arteev/dsql/app.Version=${VERSION} -X github.com/arteev/dsql/app.DateBuild=${BUILD_TIME}  -X github.com/arteev/dsql/app.GitHash=${GITHEAD}" 

default: build

lint:
	go fmt 
	go vet 
	gometalinter --deadline=15s ./...

build: lint
	go build ${LDFLAGS} -o dsql 

run: 
	go run ${LDFLAGS} main.go

zip: build	
	zip dsql-linux-$(shell arch).zip dsql

test: 
	go test -v ./...