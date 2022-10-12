APP_NAME = morning-night-dream-platform

export
POSTGRES_VERSION := 14.5
POSTGRES_IMAGE := postgres:$(POSTGRES_VERSION)-alpine
POSTGRES_HOST := localhost
POSTGRES_DB := postgres
POSTGRES_USER := postgres
POSTGRES_PASS := postgres
POSTGRES_PORT := 54321
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASS)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable
API_KEY = local

.PHONY: aqua
aqua: # export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
	@go run github.com/aquaproj/aqua-installer@latest

.PHONY: tool
tool:
	@aqua i

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: package
package:
	@go get -u

.PHONY: test
test:
	@go test ./...

.PHONY: core
core:
	@go run app/core/main.go

.PHONY: gen
gen:
	@go generate ./...

.PHONY: migrate
migrate:
	@go run cmd/migrate/main.go

.PHONY: db
db:
	@docker run --rm -d \
		-p $(POSTGRES_PORT):5432 \
		-e TZ=UTC \
		-e LANG=ja_JP.UTF-8 \
		-e POSTGRES_HOST_AUTH_METHOD=trust \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASS) \
		-e POSTGRES_INITDB_ARGS=--encoding=UTF-8 \
		--name ${APP_NAME}-db \
		$(POSTGRES_IMAGE)

.PHONY: psql
psql:
	@docker exec -it ${APP_NAME}-db psql -U postgres

.PHONY: dbstop
dbstop:
	@docker stop ${APP_NAME}-db

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o main
