#!/bin/bash

pushd ../../..
make chess.go.zip
popd
cp ../../../chess.go.zip .
