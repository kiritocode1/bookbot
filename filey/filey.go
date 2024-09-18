package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func updatePerson(p *Person) {
    p.Name = "John Doe"
    p.Age = 30
}

func main() {
    person := Person{Name: "Alice", Age: 25}
    fmt.Println("Before:", person)

    updatePerson(&person) // Pass the address of the struct

    fmt.Println("After:", person)
}