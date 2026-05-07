FROM golang:1.26-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /build/server .
ENV PORT=8080
USER nobody

CMD ["./server"]
