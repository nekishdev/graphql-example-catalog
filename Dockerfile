FROM golang:1.14.2

MAINTAINER nekishdev

ENV PROJECT_DIR=/app
ENV GOPROXY=direct
#ENV GO111MODULE=on

RUN mkdir -p $PROJECT_DIR

WORKDIR $PROJECT_DIR

COPY . .

RUN env

RUN go mod vendor

RUN go build -o ./server

EXPOSE 80

CMD ["./server"]