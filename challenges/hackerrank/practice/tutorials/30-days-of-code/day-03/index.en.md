---
title: "30 Days of Code - Day 3: Intro to Conditional Statements"
date: 2020-03-24T15:00:00-04:00
metadata:
  difficulty: Easy
  platform: HackerRank
  url: https://www.hackerrank.com/challenges/30-conditional-statements
tags:
  - challenges-easy
  - challenges-hackerrank
  - go
---

# Objective

In this challenge, we're getting started with conditional statements. Check out
the Tutorial tab for learning materials and an instructional video!

# Task

Given an integer, `n`, perform the following conditional actions:

* If `n` is odd, print `Weird`

* If `n` is even and in the inclusive range of `2` to `5`, print `Not Weird`

* If `n` is even and in the inclusive range of `6` to `20`, print `Weird`

* If `n` is even and greater than `20`, print `Not Weird`

Complete the stub code provided in your editor to print whether or not `n` is
wierd.

# Input Format

A single line containing a positive integer, `n`.

**Constraints:**

* `1 <= n <= 100`

# Output Format

Print `Weird` if the number is weird; otherwise, print `Not Weird`.

# Sample 00

{{< challenge-sample >}}

## Explanation

`n` is odd and odd numbers are weird, so we print `Weird`.

# Sample 01

{{< challenge-sample "01" >}}

## Explanation

`n > 20` and `n` is even, so it isn't weird. Thus, we print `Not Weird`.

# Solution

{{< snippet path="main.go" hl="go" foldable=true >}}

