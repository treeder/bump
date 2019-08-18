set -e

oldv=$(git describe --match "v[0-9]*" --abbrev=0 HEAD)
echo "oldv: $oldv"

newv=$(docker run --rm -v "$PWD":/app treeder/bump --input "$oldv" patch)
echo "newv: $newv"

git tag -a "v$newv" -m "version $newv"
git push --follow-tags

#