.PHONY: test validate

test:
	go test ./scripts/validate/...

validate:
	go run ./scripts/validate/ .
