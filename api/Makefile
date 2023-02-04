.PHONY: tool
tool:
	@aqua i && npm install

.PHONY: lint
lint:
	@(cd proto && buf lint)

.PHONY: format
format:
	@buf format -w

.PHONY: backend
backend:
	@buf generate --template buf.backend.gen.yaml
	@oapi-codegen -generate types -package openapi openapi.yaml > ../backend/pkg/openapi/types.gen.go
	@oapi-codegen -generate chi-server -package openapi openapi.yaml > ../backend/pkg/openapi/server.gen.go

.PHONY: frontend
frontend:
	@buf generate --template buf.frontend.gen.yaml

.PHONY: doc
doc:
	@rm -rf doc && mkdir -p doc/api && protoc --doc_out=./doc/api --doc_opt=markdown,README.md \
		proto/article/**/*.proto \
		proto/auth/**/*.proto \
		proto/health/**/*.proto
