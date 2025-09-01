FROM golang:1.24-alpine

WORKDIR /app

# Устанавливаем необходимые пакеты
RUN apk add --no-cache git bash postgresql-client

# Устанавливаем гуся
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Копируем модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем приложение
COPY . .

# Собираем бинарник
RUN go build -o my-go-app .

# Копируем entrypoint
COPY entrypoint.sh /app/entrypoint.sh
RUN chmod +x /app/entrypoint.sh
#pfgecr
ENTRYPOINT ["/app/entrypoint.sh"]