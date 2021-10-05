build:
	go build -o bin/spcli main.go
setup:
	go mod download
test:
	go test -v ./...
