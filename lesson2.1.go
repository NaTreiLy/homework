package main

import "fmt"
import "strconv"

var x float64 = 10.8
var y int = int(x)
var y1 float64 = float64(y)

var isZero bool = 4 == 0

var hello string = "Hello, World!"

var z int = 8
var g int = 4

var a bool = true
var b bool = false

var r string = "7"

var x0 float64 = 6.7

var isMen = true
var isWomen = false

var s int = 6

func main() {
	z1 := z + g
	z2 := z - g
	z3 := z * g
	z4 := z / g

	a1 := a && b
	a2 := a || b
	a3 := !b

	var r1, err = strconv.Atoi(r)
	if err != nil {
	} else {
		r2 := r1 * 2
		fmt.Println(r2)
	}

	x1 := x * x0
	x2 := x + x0
	x3 := x1 - x2

	/*которая выводит результат логической операции "отрицание" (!) для каждой из переменных. если не мужчина, то женщина?*/

	var nina = !isMen
	if nina {
		fmt.Println("Где ужин?")
	} else {
		fmt.Println("Где цветы?")
	}

	var bob = !isWomen
	if bob {
		fmt.Println("Где ужин?")
	} else {
		fmt.Println("Где цветы?")
	}

	fmt.Println(s)

	var s int = 4
	fmt.Println(s)

	fmt.Println()
	fmt.Println(x, y, y1, isZero, hello, z1, z2, z3, z4, a1, a2, a3, x3)
}
