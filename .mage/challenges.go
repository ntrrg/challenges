// +build mage

package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	challengesDir = filepath.Clean("content/challenges")
)

type Challenges mg.Namespace

func (Challenges) Default() {
	mg.SerialDeps(Challenges.Lint, Challenges.Run)
}

func (Challenges) Clean() error {
	fn := func(path string, info os.FileInfo, err error) error {
		if filepath.Base(path) != "solution" {
			return nil
		}

		return sh.Rm(path)
	}

	return filepath.Walk(challengesDir, fn)
}

func (Challenges) Lint() error {
	var files []string

	fn := func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		files = append(files, path)
		return nil
	}

	if err := filepath.Walk(challengesDir, fn); err != nil {
		return err
	}

	args := []string{"-d", "-e", "-s"}
	args = append(args, files...)
	return sh.RunV("gofmt", args...)
}

func (Challenges) Run() error {
	c := exec.Command("./run.sh")
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Dir = challengesDir

	return c.Run()
}

func init() {
	cleanDeps = append(cleanDeps, Challenges.Clean)
	lintDeps = append(lintDeps, Challenges.Lint)
}
