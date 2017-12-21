VERSION="1.6.25"
BUILD_TIME=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
GITHEAD=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X github.com/arteev/dsql/app.Version=${VERSION} -X github.com/arteev/dsql/app.DateBuild=${BUILD_TIME}  -X github.com/arteev/dsql/app.GitHash=${GITHEAD}" 

default: build

lint:
	go fmt 
	go vet 
	gometalinter --deadline=15s ./...

build: 
	go build ${LDFLAGS} -o dsql 	

cross:
	CGO_ENABLED=1 GOOS=windows GOARCH= CC=x86_64-w64-mingw32-gcc-win32 go build ${LDFLAGS} -o dsql.exe

run: 
	go run ${LDFLAGS} main.go

zip: build	
	zip dsql-linux-$(shell arch)-${VERSION}.zip dsql

zipcross: cross
	zip dsql-win64-${VERSION}.zip dsql.exe


test: 
	go test -v ./...