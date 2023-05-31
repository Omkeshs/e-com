FROM ubuntu:latest
RUN apt-get update -y

# Build stage
FROM golang:1.20.4-alpine3.18 as build
WORKDIR /app
COPY . .
# ENV GOROOT="/usr/local/go"
ENV GO111MODULE=on
ENV GOPATH="/app/order/go.mod"
RUN CGO_ENABLED=0 go build -o main cmd/main.go

# Run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/main .


EXPOSE 8080
CMD [ "/app/main" ]