package main

import "fmt"

func main() {
	var (
		firstName string
		lastName  string
		id, tests int
	)

	fmt.Scan(&firstName, &lastName, &id, &tests)
	scores := make([]int, tests)

	for i := 0; i < tests; i++ {
		fmt.Scan(&scores[i])
	}

	s := NewStudent(firstName, lastName, id, scores)
	s.printPerson()
	s.calculate()
	fmt.Printf("Grade: %v", s.calculate())
}

type Person struct {
	firstName string
	lastName  string
	id        int
}

func NewPerson(firstName, lastName string, id int) Person {
	return Person{firstName, lastName, id}
}

func (p Person) printPerson() {
	fmt.Printf("Name: %v, %v\nID: %v\n", p.lastName, p.firstName, p.id)
}

type Student struct {
	Person
	scores []int
}

func NewStudent(firstName, lastName string, id int, scores []int) Student {
	return Student{NewPerson(firstName, lastName, id), scores}
}

func (s Student) calculate() string {
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
