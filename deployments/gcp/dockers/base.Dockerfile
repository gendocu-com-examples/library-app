FROM golang:1.16-buster as builder
ENV GO111MODULE on
RUN apt-get update && apt-get install -y make git upx ca-certificates

WORKDIR /app
COPY ./backend/go.mod .
COPY ./backend/go.sum .
RUN  go mod download

COPY ./backend/cmd ./cmd
COPY ./backend/pkg ./pkg

RUN go build -o grpc-server -ldflags="-s -w" cmd/grpc/main.go && upx grpc-server
RUN go build -o grpcweb-server -ldflags="-s -w" cmd/localserver/main.go && upx grpcweb-server
