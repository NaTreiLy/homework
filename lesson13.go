package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var x1 int = 77
var y2 int = 7

type MyErrors struct {
	Message string
	File    string
	Line    int
}

func main() {

	fmt.Printf("Функция вернула %w\n", returnError1())

	err := returnError()
	if err != nil {
		fmt.Printf("Error %w\n", err)
	}

	err1 := newErrors()
	if err1 != nil {
		fmt.Printf("%w\n", err1)
	}

	fmt.Println(maibyError(x1, y2))
	fmt.Println(maibyError(y2, x1))

	f, err4 := os.Open("lesson8.go")
	if err4 != nil {
		log.Fatal(err4)
	}
	fmt.Println(f)
}
func returnError1() error {
	var err2 error
	return err2
}

func returnError() error {
	return fmt.Errorf("Истерика")
}

func newErrors() error {
	newError := errors.New("Ничего непонятно")
	return newError
}

func maibyError(x, y int) (int, error) {
	sum := x - y
	if sum < 0 {
		return 0, fmt.Errorf("Отрицательная сумма")
	}
	return sum, nil
}
