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

### Makefile targets:

- `build` - Builds binary to file path `"bin/mailer"`
- `run_tests` - Run Golang tests.
- `test_coverage` - Run Golang test converage.
- `show_coverage` - Run and displays html test coverage.
- `cleanup` - Removes temporary files.