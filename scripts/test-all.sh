#!/bin/sh -e

script_dir=$(cd `dirname $0`; pwd)
root_dir=`dirname $script_dir`

$script_dir/build-all.sh
(cd $root_dir && go test ./...)
