#!/bin/sh

tarfile="$1.tar.gz"

echo "开始打包$tarfile..."

export GOARCH=amd64
export GOOS=linux

bee pack

mv shopoms.tar.gz $tarfile
