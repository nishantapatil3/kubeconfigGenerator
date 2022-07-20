.PHONY: build

build:
	go build -o build/kubeconfigGenerator
	echo "${PWD}/build/kubeconfigGenerator"

docker-build:
	GOARCH=amd64 GOOS=linux go build -o build/kubeconfigGenerator
	docker build --tag nishantapatil3/kubeconfiggenerator:latest . -f ./Dockerfile
	docker push docker.io/nishantapatil3/kubeconfiggenerator:latest

clean:
	rm -rf build
