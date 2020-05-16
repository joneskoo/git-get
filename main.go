// Command git-get clones Git repositories with an implicitly relative URL
// and always to a path under source root regardless of working directory.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
)

const (
	// sourceRoot is the target prefix where we clone to,
	// can be overridden with environment variable GIT_GET_ROOT.
	defaultTargetPath = "~/src"

	// defaultPrefix is prefixed to implicitly relative clone URLs,
	// can be overriden with environment variable GIT_GET_PREFIX.
	defaultPrefix = "git@github.com:"
)

const usage = `git-get (URL|PROJECT/REPOSITORY)

  $ git get joneskoo/git-get                    # PROJECT/REPOSITORY
  $ git get git@github.com:joneskoo/git-get     # URL

Regardless of working directory where git get is executed, this expands to:

  $ git clone git@github.com:joneskoo/git-get ~/src/github.com/joneskoo/git-get

This allows easy cloning of repositories into an uniform directory structure.
`

func main() {
	logger := log.New(os.Stderr, "", 0)

	if len(os.Args) != 2 {
		logger.Fatalln(usage)
	}
	relativeCloneURL := os.Args[1]

	targetPath := defaultTargetPath
	if s := os.Getenv("GIT_GET_ROOT"); s != "" {
		targetPath = s
	}
	var err error
	targetPath, err = homedir.Expand(targetPath)
	if err != nil {
		logger.Fatalf("git-get: failed to expand target path: %v", err)
	}

	prefix := defaultPrefix
	if s := os.Getenv("GIT_GET_PREFIX"); s != "" {
		prefix = s
	}
	cloneURL := expand(relativeCloneURL, prefix)
	td, err := targetDir(cloneURL)
	if err != nil {
		logger.Fatalf("git-get: %v", err)
	}

	// Replace current process with git
	cmd := exec.Command("git", "clone", cloneURL, filepath.Join(targetPath, td))
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if ee, ok := err.(*exec.ExitError); ok {
		os.Exit(ee.ExitCode())
	}
	if err != nil {
		logger.Fatalf("git-get: calling git failed: %v", err)
	}
}

// expand completes the implicitly relative clone URL s of form "project/repo" into
// absolute clone URL.
//
// If s does not have the required form, it is returned unchanged.
// unmodified.
func expand(s, defaultPrefix string) string {
	if strings.Contains(s, ":") {
		return s
	}
	parts := strings.SplitN(s, "/", 2)
	if len(parts) < 2 {
		return s
	}
	return defaultPrefix + parts[0] + "/" + parts[1] + ".git"
}

// targetDir resolves the cloneURL to a relative directory path.
func targetDir(cloneURL string) (string, error) {
	cleanedCloneURL := strings.TrimSuffix(cloneURL, ".git")

	var hostname, path string

	// URLs like https:// and ssh://
	if parts := strings.SplitN(cleanedCloneURL, "://", 2); len(parts) == 2 {
		if addressParts := strings.SplitN(parts[1], "/", 2); len(addressParts) == 2 {
			hostname = addressParts[0]
			path = addressParts[1]
		} else {
			return "", fmt.Errorf(`expected path in URL, got %q`, cloneURL)
		}
		// URLs like user@hostname:project/repo
	} else if parts := strings.SplitN(cleanedCloneURL, ":", 2); len(parts) == 2 {
		hostname = parts[0]
		path = parts[1]
	} else {
		return "", fmt.Errorf(`expected PROJECT/REPO or absolute git clone URL, got %q`, cloneURL)
	}

	// ignore username
	parts := strings.Split(hostname, "@")
	hostname = parts[len(parts)-1]

	pathparts := strings.Split(path, "/")
	target := append([]string{hostname}, pathparts...)
	return strings.ToLower(filepath.Join(target...)), nil
}
