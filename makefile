.PHONY:
build:
	go build -o .bin/app.exe cmd/site/main.go

run: build
	./.bin/app.exe

test_service: 
	go test ./test/service