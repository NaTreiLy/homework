package main

import "fmt"

func main() {

	x := []int{-6, 9, 97, -87, 5}
	for index, value := range x {
		if value > 0 {
			fmt.Printf("Число №%d - положительное число\n", index)
		}
	}

	y := []int{9, 19, 756, 84, 2, 11, 7}
	plusX := make([]int, 0, 0)
	minusX := make([]int, 0, 0)
	for _, value := range y {
		if value > 10 {
			plusX = append(plusX, value)
		} else {
			minusX = append(minusX, value)
		}
	}
	//pX := len(plusX)
	//mX := len(minusX)
	fmt.Printf("В нашем слайсе %d чисел больше 10 и %d - меньше 10\n", len(plusX), len(minusX))

	for _, value := range y {
		if value < 10 {
			fmt.Printf("%d - меньше 10\n", value)
		} else if value < 20 {
			fmt.Printf("%d - меньше 20, но больше 10\n", value)
		} else if value < 30 {
			fmt.Printf("%d - меньше 30, но больше 20\n", value)
		} else {
			fmt.Printf("%d - больше 30\n", value)
		}
	}
	z := []int{1, 2, 3, 4, 5, 6, 7}
	for _, day := range z {
		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Saturday")
		case 7:
			fmt.Println("Sunday")
		}
	}

	a := []int{6, 9, 2, 1, 4, 7}
	for _, aa := range a {
		switch aa {
		case 6:
			fmt.Println("Четное")
		case 9:
			fmt.Println("Нечетное")
		case 2:
			fmt.Println("Четное")
		case 1:
			fmt.Println("Нечетное")
		case 4:
			fmt.Println("Четное")
		case 7:
			fmt.Println("Нечетное")
		}
	}

	/*b := []bool{true, false, false, true, true}
	for _, bb := range b {
		switch bb {

		}
	}*/

	c := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, cc := range c {
		switch cc {
		case 1:
			fmt.Println("Первый")
		case 2:
			fmt.Println("Второй")
		case 3:
			fmt.Println("Третий")
		case 4:
			fmt.Println("Четвертый")
		case 5:
			fmt.Println("Пятый")
		case 6:
			fmt.Println("Шестой")
		case 7:
			fmt.Println("Седьмой")
		case 8:
			fmt.Println("Восьмой")
		case 9:
			fmt.Println("Девятый")
		case 10:
			fmt.Printf("Десятый")
		}
	}
}
