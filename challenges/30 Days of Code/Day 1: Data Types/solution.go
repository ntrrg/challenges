package main

import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

  var i uint64 = 4
  var d float64 = 4.0
  var s string = "HackerRank "

  scanner := bufio.NewScanner(os.Stdin)

  scanner.Scan()
  i2, _ := strconv.ParseUint(scanner.Text(), 10, 64)
  scanner.Scan()
  d2, _ := strconv.ParseFloat(scanner.Text(), 64)
  scanner.Scan()
  s2 := scanner.Text()

  fmt.Println(i + i2)
  fmt.Printf("%.1f\n", d + d2,)
  fmt.Println(s + s2)
}