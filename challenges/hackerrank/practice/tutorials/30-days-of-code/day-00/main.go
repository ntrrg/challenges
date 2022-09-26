package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data, _ := ioutil.ReadAll(os.Stdin)
	fmt.Println("Hello, World.")
	fmt.Println(string(data))
}
