# Этап сборки
FROM golang:alpine AS builder


# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код проекта
COPY . .

# Собираем бинарный файл
# Предполагаем, что точка входа находится в cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/

# Финальный этап
FROM alpine:latest

# Устанавливаем необходимые сертификаты для HTTPS запросов
RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Копируем бинарник из этапа сборки
COPY --from=builder /app/main .

# Копируем директории с конфигами и данными (если нужны)
# Раскомментируйте, если в этих папках есть файлы, которые нужны во время выполнения
# COPY --from=builder /app/db ./db
# COPY --from=builder /app/mailer/templates ./mailer/templates

# Открываем порт (замените на ваш порт, например 8080)
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]