package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	phoneBook := make(map[string]int)

	for i := 0; i < n; i++ {
		var (
			name        string
			phoneNumber int
		)

		fmt.Scan(&name, &phoneNumber)
		phoneBook[name] = phoneNumber
	}

	for {
		var name string

		if _, err := fmt.Scan(&name); err != nil {
			break
		}

		if phone, ok := phoneBook[name]; ok {
			fmt.Printf("%v=%v\n", name, phone)
		} else {
			fmt.Println("Not found")
		}
	}
}
