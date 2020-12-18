.PHONY:
.SILENT:

build:
	go build -o ./.bin/array main.go

run: build
	./.bin/array

test:
	go test ./... -coverprofile /dev/null