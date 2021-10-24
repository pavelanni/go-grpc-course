FROM golang:1.17 as builder
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM docker.io/alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]


