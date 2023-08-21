# split into 2 steps to optimize docker image size

# stage 1: build app
FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .


# stage 2: expose builded app
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/tshare tshare

EXPOSE 3000
ENTRYPOINT ["./tshare"]