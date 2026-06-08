FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server/

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY --from=builder /app/internal/infra/database/migrations ./internal/infra/database/migrations
RUN mkdir -p file_storage
EXPOSE ${PORT:-8080}
CMD ["./server"]