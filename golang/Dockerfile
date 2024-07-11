FROM golang:1.22.4 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./src

FROM golang:1.22.4 AS runtime

WORKDIR /app

COPY --from=builder /app/main main

EXPOSE 80

ENTRYPOINT ["/app/main"]
