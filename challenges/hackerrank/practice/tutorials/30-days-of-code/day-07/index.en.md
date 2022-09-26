---
title: "30 Days of Code - Day 7: Arrays"
date: 2020-03-26T13:50:00-04:00
metadata:
  difficulty: Easy
  platform: HackerRank
  url: https://www.hackerrank.com/challenges/30-arrays
tags:
  - challenges-easy
  - challenges-hackerrank
  - go
---

# Objective

Today, we're learning about the Array data structure. Check out the Tutorial
tab for learning materials and an instructional video!

# Task

Given an array, `A`, of `N` integers, print `A`'s elements in reverse order as
a single line of space-separated numbers.

# Input Format

The first line contains an integer, `N` (the size of our array).

The second line contains `N` space-separated integers describing array `A`'s elements.

**Constraints:**

* `1 <= n <= 1000`
* `1 <= A[i] <= 10000`, where `A[i]` is the `i`th integer in the array.

# Output Format

Print the elements of array `A` in reverse order as a single line of
space-separated numbers.

# Sample

{{< challenge-sample >}}

# Solution

{{< snippet path="main.go" hl="go" foldable=true >}}

