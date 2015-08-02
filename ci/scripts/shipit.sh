#!/bin/bash

set -e

# change to root of bosh release
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $DIR/../..

set -e

version=$(cat ../version/number)
if [ -z "$version" ]; then
  echo "missing version number"
  exit 1
fi
if [[ "${release_name}X" == "X" ]]; then
  echo "missing \$release_name"
  exit 1
fi

godep restore

echo Prepare github release information
set -x
mkdir -p releases
cp ci/release_notes.md releases/notes.md
echo "${release_name} v${version}" > releases/name
echo "v${version}" > releases/tag

git config --global user.email "drnic+bot@starkandwayne.com"
git config --global user.name "Stark and Wayne CI Bot"

git merge --no-edit ${promotion_branch}

goxc -bc="linux,!arm darwin,amd64" -d=releases -pv=${version}

git add -A
git commit -m "release v${version}"
