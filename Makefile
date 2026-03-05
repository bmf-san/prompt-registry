.PHONY: test validate

test:
	go test ./scripts/

validate:
	go run ./scripts/validate.go .
