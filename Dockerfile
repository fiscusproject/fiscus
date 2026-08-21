# syntax=docker/dockerfile:1
FROM golang:1.26-alpine AS builder

WORKDIR /build
RUN apk update && apk add --no-cache make
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN make build

FROM alpine:3.22

RUN addgroup -S app && adduser -S -G app app
COPY --from=builder --chown=app:app /build/dist/server /
USER app

EXPOSE 8888

CMD ["/server"]
