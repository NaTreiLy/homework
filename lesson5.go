package main

import "fmt"

func main() {

	numbers := []int{1, 86, 95, 73, 9, 7, 0}
	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))

	tree := []string{"lemon", "cherry", "bereza"}
	tree[2] = "apple"
	fmt.Println(tree)

	food := make([]string, 5, 5)
	food = append(food, "banana", "papaya", "pomfrie", "kikiriki", "krompir")
	fmt.Println(food)

	weekend := make([]string, 0, 7)
	weekend = append(weekend, "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday")
	workDays := weekend[0:5]
	fmt.Println(workDays)

	tutu := make([]int, 0, 6)
	tutu = append(tutu, 4, 6, 3, 0, 0, 0)
	lulu := make([]int, 0, 6)
	lulu = append(lulu, 0, 0, 0, 5, 9, 3)
	copy(tutu, lulu)
	fmt.Println(tutu)
}
