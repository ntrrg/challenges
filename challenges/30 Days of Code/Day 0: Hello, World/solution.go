package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  fmt.Println("Hello, World.")
  fmt.Println(read_stdin())
}

func read_stdin() string {
  var data string
  scanner := bufio.NewScanner(os.Stdin)

  for scanner.Scan() {
    data += scanner.Text()
  }

  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, "Error:", err)
    os.Exit(1)
  }

  return data
}
