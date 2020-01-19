package main

import "fmt"

func main() {
	var T, age int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)

		p := NewPerson(age)
		p.amIOld()

		for j := 0; j < 3; j++ {
			p.yearPasses()
		}

		p.amIOld()
		fmt.Println()
	}
}

type Person struct {
	age int
}

func NewPerson(age int) Person {
	if age <= 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		age = 0
	}

	return Person{age}
}

func (p Person) amIOld() {
	var s string

	if p.age <= 12 {
		s = "You are young."
	} else if p.age <= 17 {
		s = "You are a teenager."
	} else {
		s = "You are old."
	}

	fmt.Println(s)
}

func (p *Person) yearPasses() {
	p.age++
}
