NAME=orcinus/omura

ifndef
	VERSION=0.1.3
endif

build:
		go get github.com/julienschmidt/httprouter
		go get gopkg.in/yaml.v2
		go build

docker:
	@echo "### docker build: builder image"
	docker build -t ${NAME}:${VERSION} .

push:
	docker tag ${NAME}:${VERSION} ${NAME}:latest
	docker push ${NAME}:${VERSION}
	docker push ${NAME}:latest

clean:
		cd $GOPATH/bin/omura

install:
		go install github.com/orcinustools/omura
		git clone https://github.com/orcinustools/repository.git
		cd $GOPATH
		./bin/omura

reinstall: clean install

test:
		go test ./...
