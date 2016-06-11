#!/bin/bash
echo "building from $RMANTREE/include/ri.h ..."
echo "patching 20.8 ..."
cd 20.8
./build.sh
cd ..
mkdir -p tmp
cp $RMANTREE/include/ri.h tmp/ri.h
cp 20.8/ri.patch tmp/ri.patch
cd tmp
patch -i ri.patch
cd ..

go build
./nheaders -package=ri -target=tmp/ri.h
echo "installing..."
mv out.txt tokens.go
gofmt -w ./tokens.go
mv tokens.go ..
mv out1.txt prototypes.go
gofmt -w ./prototypes.go
mv prototypes.go ..
