// Copyright 2019 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var (
		i uint64  = 4
		d float64 = 4.0
		s string  = "HackerRank "
	)

	scn := bufio.NewScanner(os.Stdin)

	// Declare second integer, double, and String variables.

	var (
		i2 uint64
		d2 float64
		s2 string
	)

	// Read and save an integer, double, and String to your variables.

	scn.Scan()
	i2, _ = strconv.ParseUint(scn.Text(), 10, 64)

	scn.Scan()
	d2, _ = strconv.ParseFloat(scn.Text(), 64)

	scn.Scan()
	s2 = scn.Text()

	// Print the sum of both integer variables on a new line.

	fmt.Println(i + i2)

	// Print the sum of the double variables on a new line.

	fmt.Printf("%.1f\n", d+d2)

	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + s2)
}
