# Build Stage
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o main main.go

# Final Stage (minimal)
FROM alpine:latest

# Install only what's needed to run the binary
RUN apk add --no-cache ca-certificates

# Copy built binary from builder
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the app port
EXPOSE 80

# Run the app
CMD ["./main"]
