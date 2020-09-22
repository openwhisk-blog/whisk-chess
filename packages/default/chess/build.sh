#!/bin/bash

set -e

pushd ../../..
echo "package main" > main/index.go
echo 'var indexHTML = `<!DOCTYPE html>' >> main/index.go
cat web/index.html | sed 's/ + "api\/default\/chess"//' >> main/index.go
echo '`;' >> main/index.go
zip -r packages/default/chess/chess.go.zip main
popd
