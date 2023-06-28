#!/bin/bash
set -euo pipefail

go get -u -t ./...
go mod tidy
go mod vendor