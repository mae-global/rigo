#!/bin/bash
echo "building from $RMANTREE/include/ri.h ..."
go build
./nheaders -package=ri -target=$RMANTREE/include/ri.h
echo "installing..."
mv out.txt tokens.go
gofmt -w ./tokens.go
mv tokens.go ..
mv out1.txt prototypes.go
gofmt -w ./prototypes.go
mv prototypes.go ..
