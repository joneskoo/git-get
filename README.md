# git-get

`git-get` is a helper that allows cloning relative URLs with a short hand

    $ git get joneskoo/git-get

Regardless of working directory where `git get` is executed, this expands to:

    $ git clone git@github.com:joneskoo/git-get ~/src/github.com/joneskoo/git-get

This allows easy cloning of repositories into an uniform
directory structure.

## Installing

    $ go get github.com/joneskoo/git-get

Make sure `git-get` is in your `PATH`; by default go get
installs to `$HOME/bin`. `git` will automatically understand
`git get` after this, but `git-get` is also valid.

## Usage

Pro tip: combine **git-get** with *CDPATH* in your shell. If you set in your `.zshrc` or `.bashrc`:

```bash
CDPATH=$HOME/src:$HOME/src/github.com:$HOME/src/github.com/joneskoo:.
```

You can use any of these commands to `cd` into `/home/joneskoo/src/github.com/joneskoo/git-get`:

    $ cd git-get
    $ cd joneskoo/git-get
    $ cd github.com/joneskoo/git-get

But not only that, you can use cd to your other favorite projects.

**WARNING: this is highly addictive and you will not be able to work without this after trying it.**
