#!/usr/bin/env bash
set -eou pipefail

##### ARGS #####
if [[ "$#" != 1 ]]; then
        echo err: no version specified
fi
version="$1"

##### NAME #####
file="go.mod"
if [[ ! -r "$file" ]]; then
	echo "error: file '${file}' not readable"
	exit 1
fi

matches=0
while read -r line; do
	if [[ "$line" =~ ^module[[:space:]].* ]]; then
		# begins with "module "
		v="$(xargs <<< "${line:7}")"

		if [[ "$matches" != 0 ]]; then
			>&2 echo "error: multiple module names found, 2nd is '${v}'"
			exit 1
		fi

        name="$(basename "${v}")"
        matches=$((matches+1))
        ((matches++))
	fi
done <go.mod

if [[ "$matches" == 0 ]]; then
	>&2 echo "error: no module name found"
	exit 1
fi

##### BUILD #####
mkdir -p "builds/$version/"{bin,release}

for arch in 386 amd64 arm arm64 mips mips64 mips64le mipsle ppc64 ppc64le riscv64 s390x; do
    export GOOS=linux
    CGO_ENABLED=0 GOARCH="$arch" go build -o "builds/$version/bin/$name-$GOOS-$arch-$version"
    # xform: flatten
    tar --xform 's:^.*/::' -czf "builds/$version/release/$name-$GOOS-$arch-$version.tgz" "builds/$version/bin/$name-$GOOS-$arch-$version"
done

sha256sum "builds/$version/bin/"*"-$version" > "builds/$version/release/sha256sum.txt"
