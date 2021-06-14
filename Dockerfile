FROM golang:1-buster AS builder

WORKDIR /src

COPY . .

RUN go build -o shufflezoommeeting .

FROM debian:buster

ENV DISCORD_TOKEN ""
ENV ZOOM_TOKEN ""
ENV ZOOM_SECRET ""
ENV ZOOM_MEETING_ID ""

COPY --from=builder /src/shufflezoommeeting /usr/local/bin/shufflezoommeeting

CMD [ shufflezoommeeting ]
