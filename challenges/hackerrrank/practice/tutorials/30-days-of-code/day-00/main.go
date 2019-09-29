// Copyright 2019 Miguel Angel Rivera Notararigo. All rights reserved.
// This source code was released under the MIT license.

package main

import (
	"fmt"
)

func main() {
	var s string

	fmt.Println("Hello, World.")
	fmt.Scanln(&s)
	fmt.Println(s)
}
