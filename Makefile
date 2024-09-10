all:
	go build .

deps:
	go mod tidy
