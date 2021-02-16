test:
	go test -v ./...
benchmark:
	go test -bench=. ./blackjack/
coverage:
	go test -v ./... -coverprofile cp.out && go tool cover -html=cp.out
vet:
	go vet .
lint:
	golint .
build:
	go build -o blackjack.gm main.go cli.go
run:
	go run main.go cli.go
web:
	gotty -w bash -r
