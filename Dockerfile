FROM golang
MAINTAINER imamdigmi <imam.digmi@gmail.com>
RUN go get github.com/julienschmidt/httprouter; \
    go get gopkg.in/yaml.v2;
ADD . /go/src/github.com/orcinustools/omura
RUN go install github.com/orcinustools/omura
ENTRYPOINT /go/bin/omura
EXPOSE 8080