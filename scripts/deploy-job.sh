#!/bin/sh

apk update && apk upgrade && apk add git openssh
eval `ssh-agent -s`
ssh-add /root/.ssh/sean_github_id_ed25519
go env -w GOPRIVATE=github.com/SeanZhenggg
go get ./...
go build -buildvcs=false -o ./main ./cmd/job && ./main