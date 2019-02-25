build:
	go install ./

linux:
	GOOS=linux GOARCH=amd64 go build -o p42 ./
	tar zcf p42.linux.amd64.tar.gz p42

darwin:
	GOOS=darwin GOARCH=amd64 go build -o p42 ./
	tar zcf p42.darwin.amd64.tar.gz p42

.PHONY: build, linux, darwin
