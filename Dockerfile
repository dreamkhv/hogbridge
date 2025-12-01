FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o HogBridge ./cmd/app/main.go

FROM mailhog/mailhog:latest

ENV GIN_MODE=release

COPY --from=builder /app/HogBridge /usr/local/bin/HogBridge

ENTRYPOINT ["sh", "-c"]
CMD ["/usr/local/bin/MailHog & /usr/local/bin/HogBridge"]
