.PHONY: all docs
all: docs speakeasy

speakeasy:
	speakeasy generate sdk --lang terraform -o . -s https://docs.api.epilot.io/opportunity.yaml

docs:
	go generate ./...

