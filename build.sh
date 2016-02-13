#!/usr/bin/env bash
SCRIPT_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

OUTPATH=$SCRIPT_DIR/out
OUT=$OUTPATH/goget-forward

GOPATH=$SCRIPT_DIR CGO_ENABLED=0 go build -a -ldflags '-s' -o $OUT main.go

docker-compose up -d

rm -rf $OUTPATH
