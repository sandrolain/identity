#!/bin/sh

go test -coverprofile=coverage.out ./src/...
go tool cover -html=coverage.out