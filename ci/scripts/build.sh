#!/bin/bash

set -e -x

ls -lat 
export GOPATH=$PWD
cd go-hello
go test
ret=$?
cd ..

if [[ ${ret} > 0 ]]; then
    exit ${ret}
fi

# craft outputs in case we use them later
mkdir dist/
cp -R Dockerfile go-hello dist/

echo "dist output:"
ls -l dist/