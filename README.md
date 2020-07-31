[![Coverage Status](https://coveralls.io/repos/github/josue/mailer/badge.svg)](https://coveralls.io/github/josue/mailer) 
[![Maintainability](https://api.codeclimate.com/v1/badges/b62ae1017a276e5752ea/maintainability)](https://codeclimate.com/github/josue/mailer/maintainability) 
[![codebeat badge](https://codebeat.co/badges/f058e4fe-7c7e-4e7f-be76-ebdc0f981424)](https://codebeat.co/projects/github-com-josue-mailer-master) 

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
