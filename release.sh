set -ex

# SET BELOW TO YOUR DOCKER HUB USERNAME
USERNAME=treeder
IMAGE=bump

# ensure working dir is clean
# if [[ -z $(git status -s) ]]
# then
#   echo "tree is clean"
# else
#   echo "tree is dirty, please commit changes before running this"
#   exit 1
# fi

# Get latest
git pull

# bump version
docker run --rm -v "$PWD":/app treeder/bump --filename VERSION "$(git log -1 --pretty=%B)"
version=`cat VERSION`
echo "version: $version"

# ./build.sh

# tag it
git add -A
git commit -m "version $version"
git tag -a "$version" -m "version $version"
git push
git push --tags

docker tag $USERNAME/$IMAGE:latest $USERNAME/$IMAGE:$version

# push it
docker push $USERNAME/$IMAGE:latest
docker push $USERNAME/$IMAGE:$version
