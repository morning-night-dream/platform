.PHONY: tool
tool:
	@aqua i && npm install

.PHONY: backend
backend:
	@oapi-codegen -generate types -package openapi openapi.yaml > ../backend/pkg/openapi/types.gen.go
	@oapi-codegen -generate chi-server -package openapi openapi.yaml > ../backend/pkg/openapi/server.gen.go
	@oapi-codegen -generate client -package openapi openapi.yaml > ../backend/pkg/openapi/client.gen.go

.PHONY: frontend
frontend:
	@docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate -i /local/openapi.yaml -g typescript-fetch -o /local/api
	@rm -r ../frontend/src/api && mv api ../frontend/src/api

.PHONY: gen
gen:
	@make frontend
	@make backend
