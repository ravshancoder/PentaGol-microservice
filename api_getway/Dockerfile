FROM golang:1.19.1-alpine3.16 as builder

WORKDIR /app

COPY . .


RUN go build -o main cmd/main.go

FROM alpine:3.16

WORKDIR /app

RUN mkdir config

COPY --from=builder /app/main .
COPY ./config/rback_model.conf ./config/rback_model.conf

EXPOSE 8080

CMD [ "/app/main" ]