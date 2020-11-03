package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{name: "parker", grade: 2}

	var s2 *student = &s1
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)

	s2.name = "peter"
	fmt.Println("student 1, name :", s1.name)
	fmt.Println("student 4, name :", s2.name)

}
