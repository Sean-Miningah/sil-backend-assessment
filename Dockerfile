FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o apiserver ./cmd/main.go

# Final stage
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/apiserver .
COPY --from=builder /app/config ./config

CMD ["./apiserver"]
