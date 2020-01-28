#!/usr/bin/env bash
export GOPATH=~/go
go install "${GOPATH}"/src/github.com/kuritayu/infra-tools/cmd/lstar
go install "${GOPATH}"/src/github.com/kuritayu/infra-tools/cmd/decotail
go install "${GOPATH}"/src/github.com/kuritayu/infra-tools/cmd/tchat
go install "${GOPATH}"/src/github.com/kuritayu/infra-tools/cmd/rapidu
go install "${GOPATH}"/src/github.com/kuritayu/infra-tools/cmd/gosf
go install "${GOPATH}"/src/github.com/kuritayu/infra-tools/cmd/xlctl
