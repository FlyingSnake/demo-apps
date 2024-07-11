FROM golang:1.22.5 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./src

FROM alpine:latest AS runtime
RUN apk add --no-cache libc6-compat

WORKDIR /app
COPY --from=builder /app/main main

EXPOSE 80

ENTRYPOINT ["/app/main"]
