#!/bin/bash

# Get the directory where the script is located
script_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
cd "$script_dir/.."

set -a
source buildScripts/.env
set +a

go run ./cmd/main.go