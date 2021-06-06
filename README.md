# cards ![build](https://github.com/abmamo/cards/workflows/build/badge.svg?branch=main)
cli blackjack game written in go

## quickstart
### terminal
to play game in terminal
```
go run main.go cli.go 
```
or 
```
make run
```
### browser
to play game in browser based terminal
```
gotty -w bash -r
```
or 
```
make web
```
then navigate to ```localhost:8080``` and start game
```
go run main.go cli.go 
```
or 
```
make run
```

## docker
if you don't have go installed locally (but have docker) you can
```
make up-cli
```
to open the repo in a local `go` container or
```
make up-web
```
to expose the docker terminal over web
