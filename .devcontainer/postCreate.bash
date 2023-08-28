#! /bin/bash

go get ./...

# install golangci-lint. binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2
golangci-lint --version

# install task for taskfile
go install github.com/go-task/task/v3/cmd/task@latest

# install swag for swagger docs generation
go install github.com/swaggo/swag/cmd/swag@latest