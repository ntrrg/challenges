---
title: "Binary Trees"
date: 2020-12-05T17:00:00-04:00
metadata:
  difficulty: Medium
  platform: The Computer Language Benchmarks Game
  url: https://benchmarksgame-team.pages.debian.net/benchmarksgame/performance/binarytrees.html
tags:
  - challenges-medium
  - challenges-benchmarksgame
  - go
---

See the [challenge description](https://benchmarksgame-team.pages.debian.net/benchmarksgame/description/binarytrees.html#binarytrees).

# Usage

```shell-session
$ go build -o solution main.go

$ echo 10 | ./solution
```

# Sample

{{< challenge-sample >}}

# Solution

This solution is not fully compliant to the challenge rules, I just wanted to
learn about profiling Go programs and tweak them to run as fast as its C and
Rust counter-parts (and indeed it is faster).

{{< snippet path="main.go" hl="go" foldable=true >}}

