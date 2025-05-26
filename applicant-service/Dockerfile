FROM golang:1.21-alpine AS builder

WORKDIR /app

# Установка git для загрузки зависимостей
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o applicant-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/applicant-service .
COPY --from=builder /app/configs ./configs
COPY --from=builder /app/frontend.html ./frontend.html

EXPOSE 8085

CMD ["./applicant-service"] 