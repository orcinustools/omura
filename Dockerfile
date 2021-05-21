FROM golang:latest

#RUN go install github.com/orcinustools/omura
RUN mkdir /app
WORKDIR /app

# Copy & build
ADD . /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /app/omura .

RUN git clone https://github.com/orcinustools/repository.git $GOPATH/bin/repository

EXPOSE 8080
ENTRYPOINT ["/app/omura"]