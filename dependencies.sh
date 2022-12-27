#!/bin/sh

git submodule update --init --recursive

go mod vendor

cd ./go-utilities

go mod vendor