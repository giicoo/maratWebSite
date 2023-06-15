.PHONY:
build:
	go build -o .bin/app cmd/site/main.go

run: build
	./.bin/app