#!/bin/bash

# set directory
script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
cd "$script_dir/.."

# to ensure that its not the old build 
rm -rf ./build/linux

export GOOS="linux"
export GOARCH="amd64"
go build -o "./build/linux/playerService" ./cmd