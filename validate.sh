#!/bin/bash
set -euo pipefail

# go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
# export PATH=$PATH:$HOME/go/bin
golangci-lint run -p metalinter,bugs,unused