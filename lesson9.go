package main

import (
	"Pn/hello"
	"fmt"
	"math"
	mr "math/rand"
	"time"
)

func main() {
	var x = time.Now().Second()
	var x2 = float64(x)
	c := math.Sqrt(x2)
	fmt.Println(c)

	var y float64 = 6.00
	var y2 = math.Pow(y, 2.00)
	fmt.Println(y2)

	m := mr.Int()
	fmt.Println(m)

	hello.Log("Hy!")
}
