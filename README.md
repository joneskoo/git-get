# git-get

`git-get` is a helper that allows cloning relative URLs with a short hand

    $ git get joneskoo/git-get

Regardless of working directory where `git get` is executed, this expands to:

    $ git clone git@github.com:joneskoo/git-get ~/src/github.com/joneskoo/git-get

This allows easy cloning of repositories into an uniform
directory structure.

[![Go](https://github.com/joneskoo/git-get/workflows/Go/badge.svg)](https://github.com/joneskoo/git-get/actions?query=workflow%3AGo)

## Installing

    $ go get -u github.com/joneskoo/git-get

Make sure `git-get` is in your `PATH`; by default go get
installs to `$HOME/bin`. `git` will automatically understand
`git get` after this, but `git-get` is also valid.

## Configuration

You can override the defaults by setting environment variables:

| Environment variable | Default           | Description                                |
| -------------------- | ----------------- | ------------------------------------------ |
| `GIT_GET_PREFIX`     | `git@github.com:` | Prefix to add to relative clone targets    |
| `GIT_GET_ROOT`       | `~/src`           | Clone destination directory root           |

## Usage

    $ git get joneskoo/git-get
    $ git get git@github.com:joneskoo/git-get
    $ git get https://github.com/joneskoo/git-get

These all clone to same directory.

Pro tip: combine **git-get** with *CDPATH* in your shell. If you set in your `.zshrc` or `.bashrc`:

```bash
CDPATH=$HOME/src:$HOME/src/github.com:$HOME/src/github.com/joneskoo:.
```

You can use any of these commands to `cd` into `/home/joneskoo/src/github.com/joneskoo/git-get`
from anywhere!

    $ cd git-get
    $ cd joneskoo/git-get
    $ cd github.com/joneskoo/git-get

But not only that, you can use cd to your other favorite projects as everything
is cloned to the same directory structure. As you can also clone with absolute
URLs, this works fine if you use this for work repositories but occasionally
clone some open source project.

**WARNING: this is highly addictive and you will not be able to work without this after trying it.**
