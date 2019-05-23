FROM golang:1.12

ARG app_env
ENV APP_ENV $app_env
ENV GO111MODULE=on

WORKDIR /go/src/github.com/maxiatanasio/golang-todo

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

EXPOSE 8080

RUN go get github.com/go-playground/justdoit

ENTRYPOINT justdoit -watch="./" -include="(.+\\.go|.+\\.c)$" -build="go build -o app ." -run="./app"