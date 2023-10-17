package main

import "fmt"

var x int = 7
var y int = 8
var p int = 16
var child string = ""
var children string = "Детство"

func main() {

	hello()

	sum(x, y)
	fmt.Println(multiplication(x, y))
	fmt.Println(twix(x, y))
	fmt.Println(quatro(x, y))
	numbers(x, y, p)
	surf(child, children)
	fmt.Println(fara())

}
func hello() {
	fmt.Println("Hi!")
}

func sum(x int, y int) {
	fmt.Println(x + y)
}

func multiplication(x int, y int) int {
	var o = x * y
	return o
}

func twix(x int, y int) (int, int) {
	var s int = x + y
	var r int = x - y
	return s, r
}

func quatro(x int, y int) int {
	var g = multiplication(x, y)
	var c = g * x * y
	return c
}

func numbers(numbers ...int) {

	var number []int
	number = append(number, x, y, p)
	fmt.Println(number)
}

func surf(child string, children string) {
	fmt.Println(child + children)
}
func fara0() int {
	var d = 11 * 5
	return d
}

func fara(...int) int {
	var result = fara0()
	var e = x*y - result
	return e
}
