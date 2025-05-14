# Build Stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main main.go

# Final Stage
FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .  

EXPOSE 80
CMD ["./main"]
