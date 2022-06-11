export tag=v1.0
root:
	export ROOT=github.com/cncamp/golang

build:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 ./homework/cmd/simple-web

release: build
	docker build -t homework/cmd/simple-web:${tag}

push: release
	docker push jrx/httpserver:v1.0