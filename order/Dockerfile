# create docker image from git repo code
FROM ubuntu:latest as builder

RUN apt-get update -y
RUN apt install git -y
RUN git clone https://github.com/Omkeshs/e-com.git

# Build stage
FROM golang:1.20.4-alpine3.18 as bin_builder
WORKDIR /app
COPY --from=builder /e-com/order/ .

ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -o main cmd/main.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=bin_builder /app/main .

CMD [ "/app/main" ]

# create docker image from local code
# # Build stage
# FROM golang:1.20.4-alpine3.18 as bin_builder
# WORKDIR /app
# # COPY . .

# ENV GO111MODULE=on
# RUN CGO_ENABLED=0 go build -o main cmd/main.go

# # Run stage
# FROM alpine:3.17
# WORKDIR /app
# COPY --from=bin_builder /app/main .

# CMD [ "/app/main" ]