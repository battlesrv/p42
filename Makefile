build:
	go install ./
	go install ./cmd/p42-cli

cli:
	go install ./cmd/p42-cli

.PHONY: build, cli
