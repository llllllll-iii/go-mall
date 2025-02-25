#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=webapi
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}