all: vet test testrace

build: deps
	go build github.com/Palen/grpc-go/...

clean:
	go clean -i github.com/Palen/grpc-go/...

deps:
	go get -d -v github.com/Palen/grpc-go/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/Palen/grpc-go/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/Palen/grpc-go/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/Palen/grpc-go/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/Palen/grpc-go/...

testdeps:
	go get -d -v -t github.com/Palen/grpc-go/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/Palen/grpc-go/...

updatedeps:
	go get -d -v -u -f github.com/Palen/grpc-go/...

updatetestdeps:
	go get -d -v -t -u -f github.com/Palen/grpc-go/...

vet: vetdeps
	./vet.sh

vetdeps:
	./vet.sh -install

.PHONY: \
	all \
	build \
	clean \
	deps \
	proto \
	test \
	testappengine \
	testappenginedeps \
	testdeps \
	testrace \
	updatedeps \
	updatetestdeps \
	vet \
	vetdeps
