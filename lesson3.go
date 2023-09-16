package main

import "fmt"

const Cifr int = 56

func main() {
	const c1 int = 7
	const Cifr int = 78
	const S string = "Лапуля"
	const F float64 = 7.8
	const B bool = true
	fmt.Println(Cifr)
	fmt.Println(c1)
	ccifr1 := c1 * Cifr
	ccifr2 := c1 + Cifr
	ccifr3 := c1 - Cifr
	ccifr4 := c1 / Cifr
	fmt.Println(ccifr1, ccifr2, ccifr3, ccifr4, S, F, B)
}
