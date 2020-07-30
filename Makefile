ALL: build

cleanup:
	rm -f cp.out

run_tests:
	go test -v ./...

test_coverage:
	go test -coverprofile cp.out ./...
	go tool cover -func=cp.out

show_coverage:
	go test -coverprofile cp.out ./...
	go tool cover -html=cp.out

build:
	mkdir -p bin
	go build -o bin/mailer ./cmd/mailer/main.go
	ls -ahl bin/mailer