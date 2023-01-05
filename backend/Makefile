export
DATABASE_URL := postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable
ENDPOINT := http://localhost:8081

.PHONY: aqua
aqua: # export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
	@go run github.com/aquaproj/aqua-installer@latest

.PHONY: tool
tool:
	@aqua i

.PHONY: doc
doc:
	@rm -rf doc && tbls doc ${DATABASE_URL} doc/databases

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: modules
modules:
	@go list -u -m all

.PHONY: renovate
renovate:
	@go get -u -t ./...

.PHONY: compile
compile:
	@go build -v ./... && go clean

.PHONY: test
test:
	@go test ./...

.PHONY: e2e
e2e:
	@go test ./e2e/... -tags e2e

.PHONY: gen
gen:
	@go generate ./...

.PHONY: migrate
migrate:
	@go run cmd/migrate/main.go
