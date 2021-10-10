help:
	@echo "list of available app commands"
	@echo
	@echo "lint        	- lint app."
	@echo "test        	- test app."
	@echo "test        	- test app."
	@echo "vet        	- vet app."
	@echo "build        	- build app."
	@echo "serve        	- serve app."
	@echo "benchmark       - benchmark app."
	@echo
	@echo "docker commands"
	@echo "up-env          - start dev container"
	@echo "up-cli          - start cli container"
	@echo "up-web          - start web container"
	@echo "exec-env        - exec into dev container"
	@echo "exec-cli        - exec into cli container"
	@echo "exec-web        - exec into web container"
	@echo "purge-env       - purge dev container"
	@echo "purge-cli       - purge cli container"
	@echo "purge-web       - purge web container"

db_container_env := cards_env
db_container_cli := cards_cli

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
	chmod +x ./scripts/go-build-all && ./scripts/go-build-all && mv cards-* builds
run:
	go run main.go cli.go

build-env:
	docker build -t ${db_container_env} -f dockerfiles/Dockerfile.cli . 
build-cli:
	docker build -t ${db_container_cli} -f ./dockerfiles/Dockerfile.cli .
up-env: build-env
	$(eval container_id = $(shell (docker ps -aqf "name=${db_container_env}")))
	$(if $(strip $(container_id)), \
		@echo "existing env container found. please run make purge-env",\
		@echo "running env container..." && docker run -it -d -v $(CURDIR):/go/src/ --name ${db_container_env} ${db_container_env}:latest /bin/bash)
	$(endif)
up-cli:
	$(eval container_id = $(shell (docker ps -aqf "name=${db_container_cli}")))
	$(if $(strip $(container_id)), \
		@echo "existing env container found. please run make purge-env",\
		@echo "running env container..." && docker run -it -d -v $(CURDIR):/go/src/ --name ${db_container_cli} ${db_container_cli}:latest /bin/bash)
	$(endif)
exec-env:
	$(eval container_id = $(shell (docker ps -aqf "name=${db_container_env}")))
	$(if $(strip $(container_id)), \
		@echo "exec into env container..." && docker exec -it ${container_id} bash,\
		@echo "env container not running.")
	$(endif)
exec-cli:
	$(eval container_id = $(shell (docker ps -aqf "name=${db_container_cli}")))
	$(if $(strip $(container_id)), \
		@echo "exec into env container..." && docker exec -it ${container_id} bash,\
		@echo "env container not running.")
	$(endif)
purge-env:
	$(eval container_id = $(shell (docker ps -aqf "name=${db_container_env}")))
	$(if $(strip $(container_id)), \
		@echo "purging env container..." && docker stop ${db_container_env} && docker rm ${db_container_env},\
		@echo "env container not running.")
	$(endif)
purge-cli:
	$(eval container_id = $(shell (docker ps -aqf "name=${db_container_cli}")))
	$(if $(strip $(container_id)), \
		@echo "purging env container..." && docker stop ${db_container_cli} && docker rm ${db_container_cli},\
		@echo "env container not running.")
	$(endif)