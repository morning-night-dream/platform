.PHONY: init
init:
	@aqua i && npm install

.PHONY: lint
lint:
	@buf lint

.PHONY: format
format:
	@buf format -w

.PHONY: backend
backend:
	@buf generate --template buf.backend.gen.yaml

.PHONY: frontend
frontend:
	@buf generate --template buf.frontend.gen.yaml

.PHONY: doc
doc:
	@protoc --doc_out=./doc/api --doc_opt=markdown,README.md article/**/*.proto auth/**/*.proto
