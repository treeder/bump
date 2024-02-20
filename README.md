# Bump

Bumps version files and other handy version tools.

## Usage

```sh
docker run --rm -it -v $PWD:/app -w /app treeder/bump [--filename FILENAME] [--input STRING] [--prerelease alpha1] [--metadata build123] [CMD]
```

You must pass in either `--filename` or `--input`.

If using `--filename`, it will overwrite the file with the new version in place, great for automation.

If using `--input`, it will write the new version to STDOUT so you can pipe that in or use it in your scripts as a variable. Example:

```sh
docker run --rm -v "$PWD":/app treeder/bump --input 1.2.3
```

Use `--index` flag to specify which found versions it should replace. 0 is first one, 2 is second one, -1 is last one, etc.

`--prerelease` will append a `-` + the prerelease value.

`--metadata` will append a `+` + the metadata value.

CMD is optional and can be one of:

* patch - default
* minor
* major

To pull the command out of your last git commit, you can add `[bump major]` or `[bump minor]` to your git commit message, then use:

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

## GitHub Action to bump version

```yaml
name: Bump version

on:
  push:
    branches: 
    - main
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - name: Bump version
      run: |
        git config --global user.email "github+actions@gmail.com"
        git config --global user.name "Actions"
        git fetch --tags
        wget -O - https://raw.githubusercontent.com/treeder/bump/master/gitbump.sh | bash
```

## GitHub Action to bump npm version AND git version 🤯

```yaml
name: Bump version

on:
  push:
    branches: 
    - main
jobs:
  bump:    
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    - name: Bump version npm style
      run: |
        git config --global user.email "github+actions@gmail.com"
        git config --global user.name "Actions"
        npm version patch
        git push --follow-tags
```

If you get this error: 

```
remote: Permission to username/repo.git denied to github-actions[bot].
fatal: unable to access 'https://github.com/username/repo/': The requested URL returned error: 403
```

Go into the repo settings -> actions -> general -> Workflow permissions and give actions read/write access.

## Extra Features

### extract

Extracts version from a string, eg: 

```sh
version=$(docker run --rm -v "$PWD":/app treeder/bump  --extract --input "`docker -v`")
```

### replace

Replace a version string with a version you provide.

```sh
docker run --rm -v "$PWD":/app treeder/bump  --filename VERSION --replace 1.2.3
```

### format

Formats a version string, eg:

```sh
docker run --rm -v "$PWD":/app treeder/bump  --extract --input "`docker -v`" --format M.m
```

Run `docker run --rm -it -v $PWD:/app -w /app treeder/bump --help` for more help.
