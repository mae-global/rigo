#!/bin/bash
echo "running regressions"

echo "> 001..."
cd 001
go run scene.go
prman -progress scene.rib
tiffdiff -dspy file -dspyfile diff.tif scene.tif reg.tif

#echo -e " \e[32mdone"
echo "done"

echo -e "all done"
