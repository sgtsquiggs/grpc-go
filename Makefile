all: vet test testrace

build: deps
	go build github.com/sgtsquiggs/grpc-go/...

clean:
	go clean -i github.com/sgtsquiggs/grpc-go/...

deps:
	go get -d -v github.com/sgtsquiggs/grpc-go/...

proto:
	@ if ! which protoc > /dev/null; then \
		echo "error: protoc not installed" >&2; \
		exit 1; \
	fi
	go generate github.com/sgtsquiggs/grpc-go/...

test: testdeps
	go test -cpu 1,4 -timeout 7m github.com/sgtsquiggs/grpc-go/...

testappengine: testappenginedeps
	goapp test -cpu 1,4 -timeout 7m github.com/sgtsquiggs/grpc-go/...

testappenginedeps:
	goapp get -d -v -t -tags 'appengine appenginevm' github.com/sgtsquiggs/grpc-go/...

testdeps:
	go get -d -v -t github.com/sgtsquiggs/grpc-go/...

testrace: testdeps
	go test -race -cpu 1,4 -timeout 7m github.com/sgtsquiggs/grpc-go/...

updatedeps:
	go get -d -v -u -f github.com/sgtsquiggs/grpc-go/...

updatetestdeps:
	go get -d -v -t -u -f github.com/sgtsquiggs/grpc-go/...

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
