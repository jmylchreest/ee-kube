#!/bin/bash
# hello-go build.sh

set -e -x

# The code is located in /hello-go
echo "List whats in the current directory"
ls -lat 

# Setup the gopath based on current directory.
export GOPATH=$PWD

# Put the binary hello-go filename in /dist
cd go-hello
go test 