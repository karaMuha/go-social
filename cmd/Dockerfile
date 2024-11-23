FROM golang:1.23-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go build -o go_social main.go

RUN chmod +x /app/go_social

# build tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/go_social /app/private_key.pem /app/

CMD [ "/app/go_social" ]