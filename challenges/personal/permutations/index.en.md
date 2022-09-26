---
title: "Permutations"
date: 2020-04-05T09:15:00-04:00
metadata:
  difficulty: Medium
tags:
  - challenges-medium
  - challenges-personal
  - go
---

Print all the permutations of the given unique natural numbers.

# Input Format

The fist line contains a single integer `N` denoting the amount of numbers.

The second line is a space-separated list of unique integers `numbers`.

**Constraints:**

* 2 <= `N` <= 10

* 0 <= `numbers[i]` <= 9

# Output Format

The output should be a line-separated list of permutations. The order is not
relevant.

# Sample 00

{{< challenge-sample >}}

# Sample 01

{{< challenge-sample "01" >}}

# Sample 02

{{< challenge-sample tc="02" open=false >}}

# Solution

{{< snippet path="main.go" hl="go" foldable=true >}}

