# Bump

Bumps version files and other handy version tools.

## Usage

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump [--filename FILENAME] [--input STRING] [CMD]
```

You must pass in either `--filename` or `--input`.

If using `--filename`, it will overwrite the file with the new version in place, great for automation.

If using `--input`, it will write the new version to STDOUT so you can pipe that in or use it in your scripts as a variable. 

CMD can be one of:

* patch - default
* minor
* major

Or to pull it out of your last git commit, you can add `[bump major]` or `[bump minor]` to your git commit message, then use:

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump --filename $version_file "$(git log -1 --pretty=%B)"
```

## Bump a Git Version

The `gitbump.sh` script will automatically bump your git tags. It will get the most recent version, bump it, 
then push the new tag.

NOTE: Ensure at least one version tag exists in the repo, eg: v0.0.0 must already exist.

Or just run this if you don't want to download the script:

```sh
wget -O - https://raw.githubusercontent.com/treeder/bump/master/gitbump.sh | bash
```

## Extra Features

### extract

Extracts version from a string, eg: 

```sh
version=$(docker run --rm -v "$PWD":/app treeder/bump  --extract --input "`docker -v`")
```

### format

Formats a version string, eg:

```sh
docker run --rm -v "$PWD":/app treeder/bump  --extract --input "`docker -v`" --format M.m
```

Run `docker run --rm -it -v $PWD:/app -w /app treeder/bump --help` for more help.
