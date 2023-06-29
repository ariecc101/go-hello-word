FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o go-hello-world

FROM alpine:3.11.3
ADD .env /app
WORKDIR /app
COPY --from=builder /app/go-hello-world /app/go-hello-world

ENTRYPOINT ["/app/go-hello-world"]
