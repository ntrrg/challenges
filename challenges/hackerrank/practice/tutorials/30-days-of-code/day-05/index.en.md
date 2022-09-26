---
title: "30 Days of Code - Day 5: Loops"
date: 2020-03-24T15:20:00-04:00
metadata:
  difficulty: Easy
  platform: HackerRank
  url: https://www.hackerrank.com/challenges/30-loops
tags:
  - challenges-easy
  - challenges-hackerrank
  - go
---

# Objective

In this challenge, we're going to use loops to help us do some simple math.
Check out the Tutorial tab to learn more.

# Task

Given an integer, `n`, print its first `10` multiples. Each multiple `n * i`
(where `1 <= i <= 10`) should be printed on a new line in the form: `n x i =
result`.

# Input Format

A single integer, `n`.

**Constraints:**

* `2 <= n <= 20`

# Output Format

Print `10` lines of output; each line `i` (where `1 <= i <= 10`) contains the
`result` of `n x i` in the form: `n x i = result`.

# Sample

{{< challenge-sample >}}

# Solution

{{< snippet path="main.go" hl="go" foldable=true >}}

