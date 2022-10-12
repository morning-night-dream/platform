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
