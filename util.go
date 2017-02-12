package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var m = sync.Mutex{}

// staticFolderPrefix is the prefix to prepend static folders with.
const staticFolderPrefix = "static"

// buildRepo at the provided Github URL.
func buildRepo(root, repo string) error {
	m.Lock()
	if err := downloadRepo(repo); err != nil {
		return err
	}
	if hasMakefile(repo) {
		m.Unlock()
		if err := build(root, repo); err != nil {
			return err
		}
		m.Lock()
	}
	m.Unlock()
	return nil
}

// downloadRepo at the provided Github URL.
func downloadRepo(repo string) error {
	if err := run(fmt.Sprintf("rm -rf %s", repoFolder(repo))); err != nil {
		return err
	}
	if err := run(fmt.Sprintf(
		"chmod -R 777 %s",
		staticFolderPrefix,
	)); err != nil {
		return err
	}
	if err := run(fmt.Sprintf(
		"git clone %s %s",
		repoURL(repo), repoFolder(repo),
	)); err != nil {
		return err
	}
	if err := run(fmt.Sprintf(
		"chmod -R 777 %s",
		staticFolderPrefix,
	)); err != nil {
		return err
	}
	return nil
}

// hasMakefile returns true if the repo has a Makefile.
func hasMakefile(repo string) bool {
	path := filepath.Join(repoFolder(repo), "Makefile")
	_, err := os.Stat(path)
	return err == nil
}

// build the repo by runnings its Makefile.
func build(root, repo string) error {
	m.Lock()
	if err := os.Chdir(repoFolder(repo)); err != nil {
		m.Unlock()
		return err
	}
	m.Unlock()
	if err := run("make"); err != nil {
		return err
	}
	m.Lock()
	if err := os.Chdir(root); err != nil {
		m.Unlock()
		return err
	}
	m.Unlock()
	return nil
}

// projectName of a repo.
func projectName(repo string) string {
	return filepath.Base(repo)
}

// repoURL returns the URL the repo can be downloaded from.
func repoURL(repo string) string {
	return "https://" + repo + ".git"
}

// repoFolder returns the path to the repo.
func repoFolder(repo string) string {
	return filepath.Join(staticFolderPrefix, repo)
}

// staticFolder returns the path to the repo's static folder.
func staticFolder(repo string) string {
	return filepath.Join(repoFolder(repo), staticFolderPrefix)
}

// run the command.
func run(cmd string) error {
	split := strings.Split(cmd, " ")
	if _, err := exec.Command(split[0], split[1:]...).Output(); err != nil {
		return err
	}
	return nil
}
