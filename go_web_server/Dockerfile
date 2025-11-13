FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 8081
CMD ["./server"]