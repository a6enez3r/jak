test:
	go test -v ./...
benchmark:
	go test -bench=. ./blackjack/
coverage:
	go test -v ./... -coverprofile cp.out && go tool cover -html=cp.out
build:
	go build -o blackjack.gm main.go
run:
	go run main.go
web:
	gotty -w bash -r
