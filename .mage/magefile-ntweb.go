// +build mage

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var (
	Default = Build

	hugoVersion  = "0.82.0"
	hugoExtended = true
	hugoPort     = "1313"
	hugoConfig   = "config.yaml"

	buildDeps []interface{}
	cleanDeps []interface{}
	lintDeps  []interface{}
)

func Build() error {
	return sh.RunV("hugo")
}

func BuildAll() error {
	mg.Deps(buildDeps...)

	return Build()
}

type BumpVersion mg.Namespace

func (BumpVersion) Hugo() error {
	fn := func(path string, info os.FileInfo, err error) error {
		switch {
		case err != nil:
			return err

		case findString(path,
			".git",
			".mage/vendor",
			"archetypes", "assets", "content", "data", "i18n", "layouts", "themes",
			"public", "resources",
		):

			return filepath.SkipDir

		case
			info.IsDir(),
			strings.HasSuffix(path, ".swp"),
			findString(path, ".mage/go.sum"):

			return nil
		}

		fd, err := os.Open(path)
		if err != nil {
			return err
		}

		defer fd.Close()

		s := bufio.NewScanner(fd)

		for s.Scan() {
			if bytes.Contains(s.Bytes(), []byte(hugoVersion)) {
				fmt.Println(path)
				break
			}
		}

		if err := s.Err(); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}

		return nil
	}

	fmt.Println("content/projects/ntweb/index.es.md")
	fmt.Println("content/projects/ntweb/index.es.md")

	return filepath.Walk(".", fn)
}

func Clean() error {
	mg.Deps(cleanDeps...)

	targets := []string{".mage/output", "public"}

	for _, target := range targets {
		if err := sh.Rm(filepath.Clean(target)); err != nil {
			return err
		}
	}

	return nil
}

func Lint() error {
	mg.SerialDeps(lintDeps...)

	files, err := filepath.Glob(".mage/*.go")
	if err != nil {
		return err
	}

	args := []string{"-d", "-e", "-s"}
	args = append(args, files...)

	return sh.RunV("gofmt", args...)
}

func Run() error {
	return sh.RunV("hugo",
		"server", "-D", "-E", "-F", "--noHTTPCache", "--i18n-warnings",
		"--disableFastRender", "--bind", "0.0.0.0", "--port", hugoPort,
		"--baseUrl", "/", "--appendPort=false",
	)
}

func findString(s string, list ...string) bool {
	for _, v := range list {
		if s == v {
			return true
		}
	}

	return false
}
