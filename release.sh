set -ex

# moved to GitHub Actions, so this file doesn't work anymore

# SET BELOW TO YOUR DOCKER HUB USERNAME
USERNAME=treeder
IMAGE=bump

# Get latest
# git pull

# bump version
# docker run --rm -v "$PWD":/app treeder/bump --filename VERSION "$(git log -1 --pretty=%B)"
# version=`cat VERSION`
# echo "version: $version"
# tag it
# git config --global user.email "treeder+github@gmail.com"
# git config --global user.name "Github Action"
# git add -A
# git commit -m "version $version"
# git tag -a "v$version" -m "version $version"
# git push --follow-tags

# # not using the file anymore, just git tags
# wget -O - https://raw.githubusercontent.com/treeder/bump/master/gitbump.sh | bash
# version=$(git tag --sort=-refname --list "v[0-9]*" | head -n 1)
# echo "new version $version"

# # docker it
# echo $DOCKER_PASS | docker login -u $DOCKER_USER --password-stdin
# docker tag $USERNAME/$IMAGE:latest $USERNAME/$IMAGE:$version
# docker push $USERNAME/$IMAGE:latest
# docker push $USERNAME/$IMAGE:$version
