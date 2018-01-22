#!/bin/sh

make -j

rm *.tgz
make tar

fileName=$(echo *.tgz)
if [ -f "$fileName" ] 
then
    sz -bye $fileName
fi
