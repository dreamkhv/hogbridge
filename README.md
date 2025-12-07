# 🚀 HogBridge

### HogBridge is a production-grade mock Mailgun API written in Go

<img src="hogbridge.png" alt="hogbridge" width="256"/>

The service accepts Mailgun-style requests (`/v3/:domain/messages`), validates the payload, automatically selects an
email formation strategy, and sends the email to MailHog via SMTP.

### Installation

```shell
git clone https://github.com/dreamkhv/hogbridge
cd hogbridge
go mod tidy

go install github.com/air-verse/air@latest # install air for hot-reload
```

### Launch

```shell
make debug # hot-reload launch via air
make run   # production launch via docker
make down  # stop and remove all containers
```

### Production

```shell
docker run --rm --name hogbridge -p 8025:8025 -p 1025:1025 -p 8080:8080 dreamkhv/hogbridge
```

```shell
SERVER_PORT: ':8080'
MAIL_HOST: localhost
MAIL_PORT: 1025
MAIL_USERNAME: null
MAIL_PASSWORD: null
```

### Example request

```shell
curl -X POST http://localhost:8080/v3/example.com/messages \
  -F from="sender@example.com" \
  -F to="receiver@example.com" \
  -F subject="Hello" \
  -F text="Hello world" \
  -F attachment=@./test.pdf
```

### License

MIT — free for personal and commercial use.