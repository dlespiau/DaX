#!/bin/bash -e

script_dir=$(cd `dirname $0`; pwd)
root_dir=`dirname $script_dir`

function test_travis
{
	echo "mode: count" > profile.cov

	for pkg in $(cat $root_dir/.test-packages.txt); do
		go test -v -covermode=count -coverprofile=profile_tmp.cov $pkg || exit 1
		[ -f profile_tmp.cov ] && tail -n +2 profile_tmp.cov >> profile.cov;
		rm -f profile_tmp.cov
	done
}

function test_local
{
	go test -v `cat $root_dir/.test-packages.txt`
}

if [ "$CI" = "true" ]; then
	test_travis
else
	test_local
fi
