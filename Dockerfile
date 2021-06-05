FROM golang:1-buster AS builder

WORKDIR /src

COPY . .

RUN go build -o sortmyvoice .

FROM debian:buster

COPY --from=builder /src/sortmyvoice /usr/local/bin/sortmyvoice

CMD [ sortmyvoice ]
