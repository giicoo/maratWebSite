.PHONY:
build:
	go build -o .bin/app cmd/site/main.go

run: build
	./.bin/app

test_service: 
	go test ./test/service

run_docker:
	docker build -t marat-web .  
	docker run -p 27017:27017 --name mongodb --net=bridge  -v mongodbdata:/data/db mongo  
	docker run -it  -p 8080:8080  --name marat marat-web 