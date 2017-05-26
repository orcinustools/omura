build:
		go get github.com/julienschmidt/httprouter
		go build

clean:
		rm omura

install: build

reinstall: clean install

test:
		go test ./...
