FROM golang:1.13.9-alpine as builder
ADD . /app
WORKDIR /app
RUN  go build main.go

FROM alpine
RUN apk add ca-certificates
WORKDIR /app
COPY --from=builder /app/configs /app/configs
COPY --from=builder /app/main /app/main
EXPOSE 8080
ENTRYPOINT ["./main"]
