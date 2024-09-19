PROJECT_NAME = ms_layout

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## init: used to initialize the Go project, tidy, docker, migration, build and deploy
.PHONY: init
init:
	go mod init gitlab.toledo24.ru/web/$(PROJECT_NAME) || true
	go mod tidy
	docker compose up -d
	go run cmd/migration/main.go
	go build -o build/package/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go
	./scripts/deploy.sh

## deploy: executing the deployment command
.PHONY: deploy
deploy:
	./scripts/deploy.sh

## fast-start: quick launch of ms_layout
.PHONY: fast-start
fast-start:
	go run cmd/$(PROJECT_NAME)/main.go

## start: build start of $(PROJECT_NAME)
.PHONY: start
start:
	go build -o build/package/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go
	build/package/$(PROJECT_NAME)

## migration-up: start the migration stage with the database
.PHONY: migration-up
migration-up:
	go run cmd/migration/main.go -action=up

## migration-down: down the migration with the database
.PHONY: migration-down
migration-down:
	go run cmd/migration/main.go -action=down

## build: build a project
.PHONY: build
build:
	go build -o build/package/$(PROJECT_NAME) cmd/$(PROJECT_NAME)/main.go

## lint: format and golangci-lint the project
.PHONY: lint
lint:
	gofmt -s -w .
	docker-compose -f docker-compose-go-lint.yml run --rm golangci-lint

## test: start test
.PHONY: test
test:
	go test -v ./...

## start-server: start systemctl server
.PHONY: start-server
start-server:
	systemctl start $(PROJECT_NAME).service

## stop-server: stop systemctl server
.PHONY: stop-server
stop-server:
	systemctl stop $(PROJECT_NAME).service

## remove-readme: delete all files README.md
.PHONY: remove-readme
remove-readme:
	find . -type f -name "*README.md" ! -path "./README.md" -exec rm -f {} +
