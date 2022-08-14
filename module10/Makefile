IMAGE ?= jrmarcco/simple-web:metric
build:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/amd64 ./cmd/simple-web

release: build
	docker build -f ./DockerFile -t $(IMAGE) .

push: release
	docker push $(IMAGE)