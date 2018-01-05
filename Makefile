VERSION="1.6.35"
BUILD_TIME=`date -u '+%Y-%m-%d_%I:%M:%S%p'`
GITHEAD=`git rev-parse HEAD`
LDFLAGS=-ldflags "-X github.com/arteev/dsql/app.Version=${VERSION} -X github.com/arteev/dsql/app.DateBuild=${BUILD_TIME}  -X github.com/arteev/dsql/app.GitHash=${GITHEAD}" 

default: build

lint:
	go fmt 
	go vet 
	gometalinter --deadline=15s ./...

dep:
	go get -v

build: 
	go build ${LDFLAGS} -o dsql 	

install:
	go install ${LDFLAGS}

cross:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc-win32 go build ${LDFLAGS} -o dsql.exe
	CGO_ENABLED=1 GOOS=windows GOARCH=386 CC=i686-w64-mingw32-gcc go build ${LDFLAGS} -o dsql-386.exe	

run: 
	go run ${LDFLAGS} main.go

zip: build	
	zip dsql-linux-$(shell arch)-${VERSION}.zip dsql

zipcross: cross
	zip dsql-win64-${VERSION}.zip dsql.exe
	zip dsql-i386-${VERSION}.zip dsql-386.exe
	rm -f dsql*.exe


test: 
	go test -v ./...

clear:
	rm -f dsql
	rm -f dsql*.exe
	rm -f dsql*.zip 