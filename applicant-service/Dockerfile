FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o applicant-service ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/applicant-service .
COPY --from=builder /app/configs ./configs

EXPOSE 8085

CMD ["./applicant-service"] 