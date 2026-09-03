#!/bin/bash

set -e

mkdir -p dist

build() {
    GOOS="$1" GOARCH="$2" go build -o "dist/ctread-$1-$2$3" .
}

build darwin arm64
build darwin amd64

build linux arm64
build linux amd64

build windows arm64 .exe
build windows amd64 .exe

build freebsd arm64
build freebsd amd64

echo "Builds complete!"
