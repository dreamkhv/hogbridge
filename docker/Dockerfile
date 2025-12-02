FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o HogBridge ./cmd/app/main.go

RUN apk --no-cache add --virtual build-dependencies \
    git \
  && mkdir -p /root/gocode \
  && export GOPATH=/root/gocode \
  && go install github.com/mailhog/MailHog@latest

FROM alpine:3.21

ENV GIN_MODE=release

RUN apk add --no-cache supervisor

COPY --from=builder /root/gocode/bin/MailHog /usr/local/bin/MailHog
COPY --from=builder /app/HogBridge /usr/local/bin/HogBridge

COPY docker/supervisord.conf /etc/supervisord.conf

ENTRYPOINT ["supervisord", "-c", "/etc/supervisord.conf"]
