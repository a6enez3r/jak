pn := jak

ifeq ($(version),)
version := 0.0.1
endif
ifeq ($(cm),)
cm := default commit message
endif
ifeq ($(branch),)
branch := main
endif
ifeq ($(env),)
env := jak:dev
endif
ifeq ($(envname),)
envname := jak
endif

# COLORS
ifneq (,$(findstring xterm,${TERM}))
	BLACK        := $(shell tput -Txterm setaf 0 || "")
	RED          := $(shell tput -Txterm setaf 1 || "")
	GREEN        := $(shell tput -Txterm setaf 2 || "")
	YELLOW       := $(shell tput -Txterm setaf 3 || "")
	LIGHTPURPLE  := $(shell tput -Txterm setaf 4 || "")
	PURPLE       := $(shell tput -Txterm setaf 5 || "")
	BLUE         := $(shell tput -Txterm setaf 6 || "")
	WHITE        := $(shell tput -Txterm setaf 7 || "")
	RESET := $(shell tput -Txterm sgr0)
else
	BLACK        := ""
	RED          := ""
	GREEN        := ""
	YELLOW       := ""
	LIGHTPURPLE  := ""
	PURPLE       := ""
	BLUE         := ""
	WHITE        := ""
	RESET        := ""
endif


TARGET_MAX_CHAR_NUM=20
## show help
help:
	@echo ''
	@echo 'usage:'
	@echo '  ${BLUE}make${RESET} ${RED}<cmd>${RESET}'
	@echo ''
	@echo 'cmds:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^## (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 3, RLENGTH); \
			printf "  ${PURPLE}%-$(TARGET_MAX_CHAR_NUM)s${RESET} ${GREEN}%s${RESET}\n", helpCommand, helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

# SCM #
## save changes locally using git
save-local:
	@echo "saving..."
	@git add .
	@git commit -m "${cm}"

## save changes to remote using git
save-remote:
	@echo "saving to remote..."
	@git push origin ${branch}

## pull changes from remote
pull-remote:
	@echo "pulling from remote..."
	@git merge origin ${branch}

## create new tag, recreate if it exists
tag:
	git tag -d ${version} || : 
	git push --delete origin ${version} || : 
	git tag -a ${version} -m "latest" 
	git push origin --tags
#######

# DEV #
## install deps [dev]
deps-dev:
	# gosec
	sudo curl -sfL https://raw.githubusercontent.com/securego/gosec/master/install.sh | sudo sh -s -- -b $(go env GOPATH)/bin v2.9.5
	# golines
	GO111MODULE=on go install github.com/segmentio/golines@latest
	# errcheck
	GO111MODULE=on go install github.com/kisielk/errcheck@latest
	# dupl
	GO111MODULE=on go install github.com/mibk/dupl@latest
	# golint
	GO111MODULE=on go install golang.org/x/lint/golint@latest
	# deps
	GO111MODULE=on go mod download
	
## cross platform build
build:
	rm -rf builds && mkdir builds && chmod +x ./scripts/go-build-all && ./scripts/go-build-all && mv ${pn}-* builds

## run package
run:
	go run main.go cli.go

## test package
test:
	go test -v ./...

## benchmark package
benchmark:
	go test -bench=. ./blackjack/

## test coverage
coverage:
	go test -v ./... -coverprofile cp.out && go tool cover -html=cp.out

## vet modules
vet:
	go vet .

## lint package
lint:
	golint .

## format package
format:
	golines main.go
	golines cli.go
	golines blackjack

## scan package for duplicate code [dupl]
scan-duplicate:
	dupl .

## scan package for errors [errcheck]
scan-errors:
	errcheck ./...

## scan package for security issues [gosec]
scan-security:
	gosec ./...
#######

# Docker #
## build docker env
build-env:
	docker build -t ${env} -f dockerfiles/Dockerfile . 

## start docker env
up-env: build-env
	$(eval cid = $(shell (docker ps -aqf "name=${envname}")))
	$(if $(strip $(cid)), \
		@echo "existing env container found. please run make purge-env",\
		@echo "running env container..." && docker run -it -d -v $(CURDIR):/go/src/ --name ${envname} ${env} /bin/bash)
	$(endif)

## exec. into docker env
exec-env:
	$(eval cid = $(shell (docker ps -aqf "name=${envname}")))
	$(if $(strip $(cid)), \
		@echo "exec into env container..." && docker exec -it ${cid} bash,\
		@echo "env container not running.")
	$(endif)

## remove docker env
purge-env:
	$(eval cid = $(shell (docker ps -aqf "name=${envname}")))
	$(if $(strip $(cid)), \
		@echo "purging env container..." && docker stop ${envname} && docker rm ${envname},\
		@echo "env container not running.")
	$(endif)

## init env + install common tools
init-env:
	apk update
	apk add --update curl
	apk add --update sudo
	apk add --update bash
	apk add --update ncurses
#######