#!/bin/bash -e

script_dir=$(cd `dirname $0`; pwd)
root_dir=`dirname $script_dir`

TEST_FLAGS="-v -race -timeout 5s"

function test_travis
{
	echo "mode: count" > profile.cov

	for pkg in $(cat $root_dir/.test-packages.txt); do
		go test $TEST_FLAGS -covermode=count -coverprofile=profile_tmp.cov $pkg
		[ -f profile_tmp.cov ] && tail -n +2 profile_tmp.cov >> profile.cov;
		rm -f profile_tmp.cov
	done
}

function test_local
{
	go test $TEST_FLAGS `cat $root_dir/.test-packages.txt`
}

if [ "$CI" = "true" ]; then
	test_travis
else
	test_local
fi
