FROM golang:1.16 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .



FROM alpine:latest
RUN apk --no-cache add ca-certificates libc6-compat
WORKDIR /root/
COPY --from=builder /app/server/ .
CMD ["./server"]