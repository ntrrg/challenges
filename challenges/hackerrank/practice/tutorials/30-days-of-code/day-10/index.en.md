---
title: "30 Days of Code - Day 10: Binary Numbers"
date: 2020-03-26T15:10:00-04:00
metadata:
  difficulty: Easy
  platform: HackerRank
  url: https://www.hackerrank.com/challenges/30-binary-numbers
tags:
  - challenges-easy
  - challenges-hackerrank
  - go
---

# Objective

Today, we're working with binary numbers. Check out the Tutorial tab for
learning materials and an instructional video!

# Task

Given a base-`10` integer,`n`, convert it to binary (base-`2`). Then find and
print the base-`10` denoting the maximum number of consecutive `1`'s in `n`'s
binary representation.

# Input Format

A single integer, `n`.

**Constraints:**

* `1 <= n <= 10^6`

# Output Format

Print a single base-`10` integer denoting the maximum number of consecutive
`1`'s in the binary representation of `n`.

# Sample 00

{{< challenge-sample >}}

## Explanation

The binary representation of `5` is `101`, so the maximum number of consecutive
`1`'s is `1`.

# Sample 01

{{< challenge-sample "01" >}}

## Explanation

The binary representation of `13` is `1101`, so the maximum number of
consecutive `1`'s is `2`.

# Solution

{{< snippet path="main.go" hl="go" foldable=true >}}

