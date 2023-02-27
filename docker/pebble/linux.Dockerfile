FROM golang:1.20.1-alpine as builder

ENV CGO_ENABLED=0

WORKDIR /pebble-src
COPY . .

RUN go build -o /go/bin/pebble ./cmd/pebble

## main
FROM alpine:3.17.2

RUN apk add --no-cache curl

COPY --from=builder /go/bin/pebble /usr/bin/pebble
COPY --from=builder /pebble-src/test/ /test/

CMD [ "/usr/bin/pebble" ]

EXPOSE 14000
EXPOSE 15000
