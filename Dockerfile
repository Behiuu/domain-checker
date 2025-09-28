# Stage 1: Build
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o domain-monitor main.go

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app

# Copy only the compiled binary
COPY --from=builder /app/domain-monitor .

RUN chmod +x /app/domain-monitor

CMD ["./domain-monitor"]

