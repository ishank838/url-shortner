FROM golang:alpine

RUN mkdir /src

ADD . /src/

WORKDIR /src

RUN go build -o main

EXPOSE 8080

CMD ["./main"]