package main

import "fmt"

type student struct {
    name string
    grade int
}

func main() {
    var s1 = student{}
    s1.name = "bruce"
    s1.grade = 2

    var s2 = student{"clark", 2}

    var s3 = student{name: "peter"}

    fmt.Println("student 1 :", s1.name)
    fmt.Println("student 2 :", s2.name)
    fmt.Println("student 3 :", s3.name)
}
