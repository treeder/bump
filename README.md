# Bump

Bumps version files and other handy version tools.

## Usage

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump [--filename FILENAME] [--input STRING] [CMD]
```

You must pass in either `--filename` or `--input`.

CMD can be one of:

* patch - default
* minor
* major

Or to pull it out of your last git commit, you can add `[bump major]` or `[bump minor]` to your git commit message, then use:

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump --filename $version_file "$(git log -1 --pretty=%B)"
```

## Extras

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
