# Build Stage
FROM golang:1.21.1-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN go build -o main main.go

# Run Stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
RUN mkdir -p /go/bin
COPY --from=builder /go/bin/goose /go/bin
COPY sql ./sql
COPY app.env .
COPY wait-for.sh .
COPY migration-and-app-start.sh .
RUN chmod +x wait-for.sh

EXPOSE 8080

ENTRYPOINT [ "/app/migration-and-app-start.sh" ]
