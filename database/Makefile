.PHONY: atlas
atlas:
	@curl -sSf https://atlasgo.sh | sh
	@atlas version

.PHONY: schema
schema:
	@atlas schema inspect -u "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable" > schema.hcl

.PHONY: migrate
migrate:
	@atlas schema apply -u "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable" --to file://schema.hcl
