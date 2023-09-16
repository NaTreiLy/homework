package main

import (
	"fmt"
	"strconv"
)

func main() {
	var u string
	var o int

	var u1, err = strconv.Atoi(u)
	if err != nil {
	} else {
		fmt.Println("Число 1")
	}
	fmt.Scan(&u1)
	fmt.Println("Число 2")
	fmt.Scan(&o)
	uo := u1 * o
	fmt.Println(uo)
}
