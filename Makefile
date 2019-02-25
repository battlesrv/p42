build:
	go install ./
	go install ./cmd/p42-cli

linux:
	GOOS=linux GOARCH=amd64 go build -o p42 ./
	GOOS=linux GOARCH=amd64 go build -o p42-cli ./cmd/p42-cli/
	tar zcf p42.linux.amd64.tar.gz p42 p42-cli

darwin:
	GOOS=darwin GOARCH=amd64 go build -o p42 ./
	GOOS=darwin GOARCH=amd64 go build -o p42-cli ./cmd/p42-cli/
	tar zcf p42.darwin.amd64.tar.gz p42 p42-cli

.PHONY: build, linux, darwin
