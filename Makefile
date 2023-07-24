PROJECTNAME := $(shell basename "$(PWD)")
OS := $(shell uname -s | awk '{print tolower($$0)}')
GOARCH := amd64

## run: execete main application in local machine
.PHONY: run
run:
	export APPENV=local; go run -race app/cmd/pfdserv/main.go

## tidy: special go mod tidy without golang database checksum(GOSUMDB) 
.PHONY: tidy
tidy:
	export GOSUMDB=off ; go mod tidy

## wiregen: generate dependency injection wire_gen.go from wire.go(+build wireinject)
.PHONY: wiregen
wiregen:
	wire ./...

## test: run go test
test:
	go clean -testcache & go test -v -race ./...

## set_private_repo_global: set a "gitdev.devops.napat.com" to be a private repo in go global environment 
set_private_repo_global:
	go env -w GOPRIVATE="gitdev.devops.napat.com/*"

## update_standard_lib: update standard library (mylab-standard-library) with GOPRIVATE option
update_standard_lib:
	GOPRIVATE=gitdev.devops.napat.com/mylab/mylab-standard-library go get gitdev.devops.napat.com/mylab/mylab-standard-library

## up: docker compose up
.PHONY: up
up:
	docker-compose up -d

## down: docker compose down
.PHONY: down
down:
	docker-compose down

## help: helper
.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Project: ["$(PROJECTNAME)"]"
	@echo " Please choose a command"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo

## gosec: run for scan code vulnerability by securego/gosec
.PHONY: gosec 
gosec: 
	gosec ./... 

## govulncheck: run for scan vulnerability package from Go vulnerability database
.PHONY: govulncheck
govulncheck: 
	govulncheck ./... 


## security: run make gosec and make govulncheck
security: gosec govulncheck
