FROM golang:latest

ENV PROJECT_DIR=/app

WORKDIR /app

RUN mkdir "/build"

COPY . .

RUN go get github.com/githubnemo/CompileDaemon && go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o /build/app" -command="/build/app"