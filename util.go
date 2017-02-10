package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// staticFolderPrefix is the prefix to prepend static folders with.
const staticFolderPrefix = "static"

// buildRepo at the provided Github URL.
func buildRepo(repo string) error {
	if err := downloadRepo(repo); err != nil {
		return err
	}
	if hasMakefile(repo) {
		if err := build(repo); err != nil {
			return err
		}
	}
	return nil
}

// downloadRepo at the provided Github URL.
func downloadRepo(repo string) error {
	if err := run(fmt.Sprintf("rm -rf %s", repoFolder(repo))); err != nil {
		return err
	}
	if err := run(fmt.Sprintf(
		"git clone %s %s",
		repoURL(repo), repoFolder(repo),
	)); err != nil {
		return err
	}
	return nil
}

// hasMakefile returns true if the repo has a Makefile..
func hasMakefile(repo string) bool {
	path := filepath.Join(repoFolder(repo), "Makefile")
	_, err := os.Stat(path)
	return err == nil
}

// build the repo by runnings its Makefile.
func build(repo string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err = os.Chdir(repoFolder(repo)); err != nil {
		return err
	}
	if err = run("make"); err != nil {
		return err
	}
	if err = os.Chdir(wd); err != nil {
		return err
	}
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
