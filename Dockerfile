FROM golang:1-buster AS builder

WORKDIR /src

COPY . .

RUN go build -o app ./cmd/shuffle

FROM debian:buster

RUN apt-get update \
  && apt-get install -y --no-install-recommends \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /src/app /usr/local/bin/app

CMD [ app ]
