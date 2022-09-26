---
title: "30 Days of Code - Day 9: Recursion"
date: 2020-03-26T14:40:00-04:00
metadata:
  difficulty: Easy
  platform: HackerRank
  url: https://www.hackerrank.com/challenges/30-recursion
tags:
  - challenges-easy
  - challenges-hackerrank
  - go
math: true
---

# Objective

Today, we're learning and practicing an algorithmic concept called *Recursion*.
Check out the Tutorial tab for learning materials and an instructional video!

# Task

Write a factorial function that takes a positive integer, `N` as a parameter
and prints the result of `N!` (`N` factorial).

**Note:** If you fail to use recursion or fail to name your recursive function
*factorial* or *Factorial*, you will get a score of `0`.

# Input Format

A single integer, `N` (the argument to pass to *factorial*).

**Constraints:**

* `2 <= n <= 12`

* Your submission must contain a recursive function named *factorial*.

# Output Format

Print a single integer denoting `N!`.

# Sample

{{< challenge-sample >}}

## Explanation

Consider the following steps:

1. `factorial(3) = 3 * factoria(2)`
2. `factorial(2) = 2 * factoria(1)`
3. `factorial(1) = 1

From steps `2` and `3`, we can say `factorial(2) = 2 * 1 = 2`; then when we
apply the value from `factorial(2)` to step `1`, we get `factorial(3) = 3 * 2 *
1 = 6`. Thus, we print `6` as our answer.

# Solution

{{< snippet path="main.go" hl="go" foldable=true >}}

