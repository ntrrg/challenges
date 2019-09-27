// Copyright 2019 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World.")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Println(s.Text())
}
