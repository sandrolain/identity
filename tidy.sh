#!/bin/sh

go mod tidy
go mod vendor

cd ./go-utilities

go mod tidy
go mod vendor