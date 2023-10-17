package main

import "fmt"

func main() {

	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 21; i += 2 {

		fmt.Println(i)
	}

	for i := 1; i < 21; i++ {
		if v := i % 2; v != 0 {
			continue
		}
		fmt.Println(i)
	}

	for i := 1; i < 21; i++ {
		if v := i % 2; v == 0 {
			continue
		}
		fmt.Println(i)
	}

	var sum = 0
	for i := 0; i < 100; i++ {
		sum += i
		if sum > 50 {
			break
		}
		fmt.Println(i)
	}

	var franceVerbe [8]string = [8]string{"Preparer", "Aller", "Avoir", "Arriver", "Etre", "Vouloir", "Pouvoir", "Rester"}
	for i, v := range franceVerbe {
		fmt.Printf("Индекс: %d - Значение: %s\n", i, v)
	}

	xs := make([]int, 0, 0)
	xs = append(xs, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	for i, v := range xs {
		fmt.Printf("Индекс: %d, Значение: %d\n", i, v)
	}

	var conjugationVerbs = map[string]string{
		"I":   "am",
		"You": "are",
		"She": "is",
	}
	for i, v := range conjugationVerbs {
		fmt.Printf("Ключ: %s - Значение: %s\n", i, v)
	}

	var london string = "London is a capital of Great Britan"
	for i := 0; i < len(london); i++ {
		fmt.Println(string(london[i]))
	}
}
