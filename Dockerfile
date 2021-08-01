FROM golang:1-buster AS builder

WORKDIR /src

COPY . .

RUN go build -o app ./cmd/shuffle

FROM debian:buster

COPY --from=builder /src/app /usr/local/bin/app

CMD [ app ]
