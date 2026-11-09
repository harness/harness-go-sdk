TEST?=$$(go list ./... | grep -v 'vendor' | grep -v 'examples')
default: testacc

build:
	go build -o ${BINARY}

test: 
	go test -short -coverprofile=coverage.out -covermode=atomic $(TEST)

prep-release:
	./utils/build.sh preprelease
