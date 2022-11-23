export
DATABASE_URL := postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable

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

.PHONY: lib
lib:
	@go list -u -m all && go get -u ./...

.PHONY: test
test:
	@go test ./...

.PHONY: gen
gen:
	@go generate ./... && go generate ./...

.PHONY: ent
ent:
	@mkdir tmp && \
	cp app/db/ent/generate.go tmp/ && \
	cp -r app/db/ent/schema/ tmp/schema/ && \
	rm -rf app/db/ent/ && \
	mkdir -p app/db/ent && \
	cp -r tmp/ app/db/ent/ && \
	rm -rf tmp

.PHONY: migrate
migrate:
	@go run cmd/migrate/main.go
