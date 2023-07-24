ARG  GO_VERSION=golang:1.17.5-alpine3.15

FROM  ${GO_VERSION}

RUN apk add git
RUN apk add openssh

ENV TZ=Asia/Bangkok

ARG APP_NAME
ENV APP_NAME=${APP_NAME}

WORKDIR /app

COPY . .

ENV GONOSUMDB gitdev.inno.ktb/*,gitdev.devops.krungthai.com/*
ENV GOPROXY https://artifact.devops.krungthai.com/repository/goproxy,https://proxy.golang.org,direct

COPY config.yaml config.yaml

# Using go mod.
RUN go mod download
RUN go mod verify
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT /go/bin/CompileDaemon -directory="." --build="go build -o ./tmp/app main.go" --command="./tmp/app"