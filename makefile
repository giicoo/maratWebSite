.PHONY:
build:
	go build -o .bin/app cmd/site/main.go

run: build
	./.bin/app

test_service: 
	go test ./test/service
