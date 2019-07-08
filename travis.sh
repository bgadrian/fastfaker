#!/usr/bin/env bash

#go 1.0 fail: cannot use test profile flag with multiple packages
[[ -v COVERAGE ]] && go test -coverprofile=coverage.txt -covermode=atomic ./... && golangci-lint run || go test ./...