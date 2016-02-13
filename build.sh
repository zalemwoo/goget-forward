#!/usr/bin/env bash
SCRIPT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

OUTPATH=$SCRIPT_DIR/out
OUT=$OUTPATH/go-get-forward

GOPATH=$SCRIPT_DIR CGO_ENABLED=0 go build -a -ldflags '-s' -o $OUT main.go

docker build -t goget-forward -f Dockerfile.forward .
docker build -t goget -f Dockerfile.goget .

rm -rf $OUTPATH
