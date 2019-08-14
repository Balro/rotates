#!/usr/bin/env bash

BIN=`dirname $0`
BIN=`cd -P $BIN;pwd`

DIST=$BIN/rotates-bin

sh $BIN/deps/get_deps.sh
if [ $? -ne 0 ]; then
    echo "get deps failed, please check."
    exit 1
fi

rm -rf $DIST
mkdir -p $DIST $DIST/bin $DIST/conf $DIST/logs

export GOPATH=$BIN/deps:$BIN/main

echo $GOPATH
GOOS=linux GOARCH=amd64 go build -o $DIST/lib/rotates rotates
if [ $? -ne 0 ]; then
    echo "go build failed, please check."
    exit 1
fi

cp $BIN/conf/config.toml.template $DIST/conf/
cp $BIN/scripts/rotates.sh $DIST/bin/

rm -f rotates.tar.gz
tar zcvf rotates.tar.gz rotates-bin
rm -rf rotates-bin
