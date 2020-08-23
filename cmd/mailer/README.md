## Mailer

Small library to send emails via Gmail SMTP server.

### 1. Create your secure [Google App Password](https://support.google.com/accounts/answer/185833?hl=en).

### 2. Configure your SMTP credentials via environment variable:
```
export MAILER_USERNAME=.........
export MAILER_PASSWORD=.........
```

### 3. Run example using `go run`
```
go run main.go -from "USER@gmail.com" -to "USER@gmail.com" \
    -subject "test mailer" \
    -body "hello from command line"
```

----

See root project [README](../README.md) for compiling and running.
