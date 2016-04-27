#! /usr/bin/env sh
echo install
mkdir -p output/strict
go env
go get -t -v ./...