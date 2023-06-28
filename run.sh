#!/bin/bash

DIR="$(cd "/Users/lucas/dotfiles/bin" >/dev/null 2>&1 && pwd)"

build() {
    go install
}

tests() {
    go test ./...
}

execute() {
    echo "execute not implemented yet"
    exit 1
}

"$@"
