.PHONY: build test vet lint fmt clean

build:
	go build ./...

test:
	go test ./...

test-v:
	go test -v ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

clean:
	go clean ./...

check: vet test
