test:
	go test -v ./...
build:
	go build -o blackjack main.go
run:
	go run main.go
web:
	gotty -w bash -r