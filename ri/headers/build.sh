#!/bin/bash
echo "building from $RMANTREE/include/ri.h..."
go run ./gen.go -package=ri -target=$RMANTREE/include/ri.h
gofmt -w ./out.txt
echo "installing..."
mv out.txt ../tokens.go

