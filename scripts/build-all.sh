#!/bin/sh -e

script_dir=$(cd `dirname $0`; pwd)
root_dir=`dirname $script_dir`

for e in `ls $root_dir/examples`; do
	echo "Building $e"
	cd $root_dir/examples/$e
	go build
done
