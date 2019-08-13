#!/usr/bin/env bash

BIN=`dirname $0`
BIN=`cd -P $BIN;pwd`

DIST=$BIN/rotate-bin

rm -rf $DIST
mkdir -p $DIST $DIST/bin $DIST/conf $DIST/logs

export GOPATH=$BIN/deps:$BIN/main

echo $GOPATH
GOOS=linux GOARCH=amd64 go build -o $DIST/lib/rotate rotate

cp $BIN/conf/config.toml.template $DIST/conf/
cp $BIN/scripts/rotate.sh $DIST/bin/

tar zcvf rotate.tar.gz rotate-bin
rm -rf rotate-bin
