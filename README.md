# 🚀 HogBridge

HogBridge is a production-grade mock Mailgun API written in Go.
The service accepts Mailgun-style requests (`/v3/:domain/messages`), validates the payload, automatically selects an
email formation strategy, and sends the email to MailHog via SMTP.

### Installation

```bash
git clone https://github.com/dreamkhv/hogbridge
cd hogbridge
go mod tidy

go install github.com/air-verse/air@latest # install air for hot-reload
```

### Launch

```shell
make debug # hot-reload launch via air
make run   # production laucnh via docker
make down  # stop and remove all containers
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