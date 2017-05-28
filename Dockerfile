FROM golang:1.7-alpine
MAINTAINER imamdigmi <imam.digmi@gmail.com>
RUN apk add --update git
RUN go get github.com/julienschmidt/httprouter; \
    go get gopkg.in/yaml.v2;
ADD . /go/src/github.com/orcinustools/omura
RUN go install github.com/orcinustools/omura
RUN cd /go/src/github.com/orcinustools/omura; git clone https://github.com/orcinustools/repository.git
ENTRYPOINT /go/bin/omura
EXPOSE 8080
