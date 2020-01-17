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
	Default  = Run
	goFiles  = getGoFiles()
	binFiles = getBinFiles()
)

func Clean() error {
	for _, p := range binFiles {
		if err := sh.Rm(p); err != nil {
			return err
		}
	}

	return nil
}

func Run() error {
	c := exec.Command("./run.sh")
	c.Stdout = os.Stdout

	return c.Run()
}

// Development

func CI() {
	mg.SerialDeps(Run, Lint)
}

func Format() error {
	args := []string{"-s", "-w", "-l"}
	args = append(args, goFiles...)
	return sh.RunV("gofmt", args...)
}

func Lint() error {
	args := []string{"-d", "-e", "-s"}
	args = append(args, goFiles...)
	return sh.RunV("gofmt", args...)
}

// Helpers

func getBinFiles() []string {
	var binFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if filepath.Base(path) != "solution" {
			return nil
		}

		binFiles = append(binFiles, path)
		return nil
	})

	return binFiles
}

func getGoFiles() []string {
	var goFiles []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		goFiles = append(goFiles, path)
		return nil
	})

	return goFiles
}
