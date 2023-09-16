package main

import "fmt"

func main() {

	var Cifr [5]int
	Cifr[0] = 6
	Cifr[1] = 3
	Cifr[2] = 5
	Cifr[3] = 8
	Cifr[4] = 94
	fmt.Println(Cifr[0])
	fmt.Println(Cifr[2])
	fmt.Println(Cifr[1])
	fmt.Println(Cifr[4])
	fmt.Println(Cifr[3])

	var girlfriends [3]string
	girlfriends[0] = "Nika"
	girlfriends[1] = "Sasha"
	girlfriends[2] = "Valentina Mihalna"
	fmt.Println(girlfriends)
	girlfriends[1] = "Kissa"
	fmt.Println(girlfriends)

	var chisla [5]float64
	chisla[0] = 6.7
	chisla[1] = 75.9
	chisla[2] = 8.45
	chisla[3] = 0.86
	chisla[4] = 70.56
	fmt.Println(chisla)
	var longMassiv = len(chisla)
	fmt.Println(longMassiv)

	var boyfriends [3]string
	boyfriends[0] = "Kolya"
	boyfriends[1] = "Bob"
	boyfriends[2] = "Viktor"

	var witch [2][3]string
	witch[0] = girlfriends
	witch[1] = boyfriends
	fmt.Println(witch)

	var animals [4]string
	animals[0] = "bird"
	animals[1] = "cat"
	animals[2] = "dog"
	animals[3] = "hippopotam"
	fmt.Println(animals[0], animals[3])
}
