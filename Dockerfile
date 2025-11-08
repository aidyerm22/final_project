# используем официальный образ Go как builder 
FROM golang:1.24-alpine AS builder

# устанавливаем рабочую директорию внутри контейнера 
WORKDIR /app

# копируем go.mod, чтобы скачать зависимости отдельно(для кэширования)
COPY go.mod ./
RUN go mod download

# копируем весь исходный код проекта в контейнер
COPY . .

# собираем приложение для Linux
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/main ./cmd/app

# минимальный контейнер 
FROM alpine:latest 

WORKDIR /root/

COPY --from=builder /app/main .

# прокидываем порт 8080 для приложения 
EXPOSE 8080

# команда запуска приложения при старте контейнера
CMD ["./main"]