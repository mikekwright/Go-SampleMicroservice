FROM golang:1.19.1-alpine3.16 AS builder
WORKDIR /SampleGoService
COPY . /SampleGoService
RUN mkdir -p /output && \
    go build -o /output/app ./src


FROM alpine:3.16
WORKDIR /
COPY --from=builder /output/app /app
ENTRYPOINT "./app"
