build:
		go get github.com/julienschmidt/httprouter
		go get gopkg.in/yaml.v2
		go build

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
