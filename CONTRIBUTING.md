# Contributing Guide

## Requirements

[Git]: https://git-scm.com/
[Go]: https://golang.org/dl/

* [Git][]
* [Go][] >= 1.16

**Optional:**

[GolangCI Lint]: https://github.com/golangci/golangci-lint/releases
[GNU Make]: https://www.gnu.org/software/make/
[reflex]: https://github.com/cespare/reflex
[swag]: https://github.com/swaggo/swag

* [GolangCI Lint][] >= 1.42
* [GNU Make][] >= 4.3 (build tool)
* [reflex][] >= 0.2 (filesystem watching)
* [swag][] >= 1.7.0 (OpenAPI generator)

## Guidelines

* **Git commit messages:** <https://chris.beams.io/posts/git-commit/>;
  additionally any commit must be scoped to the package where changes were
  made, which is prefixing the message with the package name, e.g.
  `cmd/api: Do something`.

* **Git branching model:** <https://guides.github.com/introduction/flow/>.

* **Version number bumping:** <https://semver.org/>.

* **Changelog format:** <http://keepachangelog.com/>.

* **Go code guidelines:** <https://golang.org/doc/effective_go.html>, <https://github.com/golang/go/wiki/CodeReviewComments>.

## Instructions

1. Create a new branch with a short name that describes the changes that you
   intend to do. If you don't have permissions to create branches, fork the
   project and do the same in your forked copy.

2. Do any change you need to do and add the respective tests.

3. **(Optional)** Run `make ci-race` (or `make ci` if your platform doesn't
   support the Go's race conditions detector) in the project root folder to
   verify that everything is working.

4. Create a [pull request](https://github.com/CapiUp/apis/compare) to the
   `main` branch.

