export
DATABASE_URL := postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable
CORE_ENDPOINT := http://localhost:8081
GATEWAY_ENDPOINT := http://localhost:8082
ENV := local
API_KEY := local

.PHONY: aqua
aqua: # export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
	@go run github.com/aquaproj/aqua-installer@latest

.PHONY: tool
tool:
	@aqua i

.PHONY: doc
doc:
	@rm -rf doc
	@tbls doc ${DATABASE_URL} doc/database
	@mkdir -p doc/proto
	@protoc --doc_out=./doc/proto --doc_opt=markdown,README.md \
		proto/article/**/*.proto \
		proto/auth/**/*.proto \
		proto/health/**/*.proto

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
	@go test ./internal/...

.PHONY: e2e
e2e:
	@go test ./e2e/... -tags e2e -count=1

.PHONY: gen
gen:
	@go generate ./... && go mod tidy

.PHONY: connect
connect:
	@buf generate --template ./proto/buf.gen.yaml

.PHONY: bufmt
bufmt:
	@buf format -w

.PHONY: buflint
buflint:
	@(cd proto && buf lint)
