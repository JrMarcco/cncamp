build:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 ./homework/cmd/simple-web

release: build
	docker build -f ./DockerFile -t jrx/simple-web:v1.0.0 ./homework

push: release
	docker push jrx/simple-web:v1.0.0