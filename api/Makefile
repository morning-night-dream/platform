.PHONY: init
init:
	@aqua i && npm install

.PHONY: lint
lint:
	@buf lint

.PHONY: format
format:
	@buf format -w

.PHONY: gen
gen:
	@buf generate
