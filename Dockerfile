# Stage 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/api

# Stage 2: Base Image
FROM alpine:latest

RUN adduser -D myuser
USER myuser

WORKDIR /home/myuser

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]