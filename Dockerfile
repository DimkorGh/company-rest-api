FROM golang:1.19.5-alpine3.17 as builder

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -tags musl -o main ./cmd/main

FROM alpine:3.17.1

WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8090
CMD ["./main"]