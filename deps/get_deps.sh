#!/usr/bin/env bash

DEPS_DIR=`dirname $0`
DEPS_DIR=`cd -P $DEPS_DIR; pwd`

mkdir -p $DEPS_DIR/src

deps=(
    "github.com/aliyun/alibaba-cloud-sdk-go/sdk"
    "github.com/BurntSushi/toml"
)

for d in ${deps[@]}; do
    g=`echo $d | sed 's/\(github.com\/[a-zA-Z_-]*\/[a-zA-Z_-]*\).*/\1/'`
    url="https://$g $DEPS_DIR/src/$g"
    dest="$DEPS_DIR/src/$g"

    if [ -e "$dest" ]; then
        cd $dest && git pull
    else
        git clone -v https://$g $DEPS_DIR/src/$g
    fi
done
