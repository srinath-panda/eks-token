CWD := $(shell pwd)

build:
	@env GOOS=darwin GOARCH=amd64 go build -o ./bin/eks-token-darwin-amd64
	@env GOOS=darwin GOARCH=arm64 go build -o ./bin/eks-token-darwin-arm64
	@env GOOS=linux GOARCH=amd64 go build -o ./bin/eks-token-linux-amd64
	@env GOOS=linux GOARCH=arm64 go build -o ./bin/eks-token-linux-arm64

