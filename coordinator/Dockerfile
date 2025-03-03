FROM golang:1.23.6 AS builder

RUN apt-get update && apt-get install -y ca-certificates git-core ssh

WORKDIR /opt/coordinator

COPY . .

ENV TZ=Europe/Moscow

# Установка зависимостей
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# & Сборка
RUN go build -o app ./cmd/app/main.go

# Сделать файл исполняемым
RUN chmod +x app

# Старт приложения
ENTRYPOINT ["./app"]
