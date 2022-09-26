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

		i2 uint64
		d2 float64
		s2 string
	)

	scn := bufio.NewScanner(os.Stdin)

	scn.Scan()
	i2, _ = strconv.ParseUint(scn.Text(), 10, 64)
	fmt.Println(i + i2)

	scn.Scan()
	d2, _ = strconv.ParseFloat(scn.Text(), 64)
	fmt.Printf("%.1f\n", d+d2)

	scn.Scan()
	s2 = scn.Text()
	fmt.Println(s + s2)
}
