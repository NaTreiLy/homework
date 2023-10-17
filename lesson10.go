package main

import (
	"fmt"
)

type Book struct {
	Title string
	Autor string
	Pages int
}

type Person struct {
	Name      string
	Age       int
	IsMarried bool
}

type Employee struct {
	Name string
	Work Job
}
type Job struct {
	Title  string
	Salary int
}

func main() {

	book1 := Book{
		"Silmarion",
		"Tolkin",
		1084,
	}

	book2 := Book{
		"128 days",
		"Sipki",
		98,
	}

	book3 := Book{
		"Alice in Wonderland",
		"Kerrol",
		278,
	}

	person1 := Person{
		"Karina",
		24,
		false,
	}

	person2 := Person{
		"Alisa",
		2,
		false,
	}

	person3 := Person{
		"Karina",
		24,
		true,
	}

	employee1 := Employee{
		Name: "Lex",
		Work: Job{
			Title:  "Cat",
			Salary: 1,
		},
	}

	persons := person1 == person3

	fmt.Println(book1, person1, person2)
	person2.Age = 3
	fmt.Println(person2, employee1, persons)

	fmt.Println(doubleBook(book1, book2))
	fmt.Println(doubleBook(book3, book2))

}
func doubleBook(x Book, y Book) Book {
	if x.Pages < y.Pages {
		return y
	} else {
		return x
	}
}
