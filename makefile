build:
	go build -v .

test: build
	go test -v ./...

