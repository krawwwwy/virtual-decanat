FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o auth-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/auth-service .
COPY --from=builder /app/configs ./configs

EXPOSE 8081

CMD ["./auth-service"] 