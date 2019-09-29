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
	Default  = Build
	goFiles  = getGoFiles()
	binFiles = getBinFiles()
)

func Build() error {
	for _, file := range goFiles {
		if filepath.Base(file) != "main.go" {
			continue
		}

		c := exec.Command("go", "build", "-o", "solution", "main.go")
		c.Dir = filepath.Dir(file)
		c.Stdout = os.Stdout

		if err := c.Run(); err != nil {
			return err
		}
	}

	return nil
}

func Clean() error {
	for _, p := range binFiles {
		if err := sh.Rm(p); err != nil {
			return err
		}
	}

	return nil
}

// Development

func CA() error {
	return sh.RunV("golangci-lint", "run")
}

func CI() {
	mg.SerialDeps(Lint, CA, Build)
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
