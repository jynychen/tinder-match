FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal
COPY model ./model
RUN go build -o bin/server cmd/server/main.go


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/server .

ENV SERVE_ADDR=0.0.0.0:8080
ENTRYPOINT ["./server"]