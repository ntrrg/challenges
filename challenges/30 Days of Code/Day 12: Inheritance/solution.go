// This challenge is intended for object oriented languages, but this file aims
// to demonstrate that it is easy (and almost the same) to use the composition
// approach in Go.

package main

import "fmt"

func main() {
	var (
		firstName, lastName string
		id, tests           int
	)

	fmt.Scan(&firstName, &lastName, &id, &tests)
	scores := []int{}

	for i := 0; i < tests; i++ {
		var score int
		fmt.Scan(&score)

		scores = append(scores, score)
	}

	s := NewStudent(firstName, lastName, id, scores)
	s.printPerson()
	s.calculate()
	fmt.Printf("Grade: %v", s.calculate())
}

type Person struct {
	firstName, lastName string
	id                  int
}

func NewPerson(firstName, lastName string, id int) *Person {
	p := new(Person)
	p.firstName = firstName
	p.lastName = lastName
	p.id = id

	return p
}

func (p *Person) printPerson() {
	fmt.Printf("Name: %v, %v\nID: %v\n", p.lastName, p.firstName, p.id)
}

type Student struct {
	Person
	scores []int
}

func NewStudent(firstName, lastName string, id int, scores []int) *Student {
	p := new(Student)
	p.firstName = firstName
	p.lastName = lastName
	p.id = id
	p.scores = scores

	return p
}

func (s *Student) calculate() string {
	avg := average(s.scores...)

	switch {
	case avg >= 90:
		return "O"
	case avg >= 80:
		return "E"
	case avg >= 70:
		return "A"
	case avg >= 55:
		return "P"
	case avg >= 40:
		return "D"
	default:
		return "T"
	}
}

func average(numbers ...int) (result int) {
	for _, n := range numbers {
		result += n
	}

	return result / len(numbers)
}
