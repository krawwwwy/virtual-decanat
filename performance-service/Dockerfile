FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY frontend.html .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o performance-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/performance-service .
COPY --from=builder /app/frontend.html .
COPY --from=builder /app/configs ./configs

EXPOSE 8084

CMD ["./performance-service"] 