build:
		go build

clean:

install: build

reinstall: clean install

test:
		go test ./...
