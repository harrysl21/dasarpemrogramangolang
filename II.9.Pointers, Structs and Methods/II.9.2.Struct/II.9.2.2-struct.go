package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
    var s1 student
    s1.name = "bruce wayne"
    s1.grade = 5

    fmt.Println("name  :", s1.name)
    fmt.Println("grade :", s1.grade)
}

