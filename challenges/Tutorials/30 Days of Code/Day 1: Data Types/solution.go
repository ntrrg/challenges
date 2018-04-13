package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scn := bufio.NewScanner(os.Stdin)

	scn.Scan()
	i2, _ := strconv.ParseUint(scn.Text(), 10, 64)
	scn.Scan()
	d2, _ := strconv.ParseFloat(scn.Text(), 64)
	scn.Scan()
	s2 := scn.Text()

	fmt.Println(i + i2)
	fmt.Printf("%.1f\n", d+d2)
	fmt.Println(s + s2)
}
