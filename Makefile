name = version
version = `cat ./version`

build:
	go build -o release/$(name)-$(version) *.go

test:
	go test -v ./...