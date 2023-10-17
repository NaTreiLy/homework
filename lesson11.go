package main

import "fmt"

type Rectangle struct {
	Weight int
	Height int
}

type Seconds struct {
	Seconds int
}

type Persons struct {
	Name string
	Live Address
}

type Address struct {
	Street  string
	City    string
	Country string
}

func main() {

	rectangle1 := Rectangle{
		5,
		5,
	}

	rectangle2 := Rectangle{
		6,
		8,
	}

	rectangle3 := Rectangle{
		9,
		2,
	}

	seconds1 := Seconds{
		78,
	}

	seconds2 := Seconds{
		98647,
	}
	live1 := Address{
		"Culture street",
		"Saint-Petersburg",
		"Russia",
	}
	persons1 := Persons{
		"Valdemar",
		live1,
	}

	fmt.Println(rectangle1.Area(), rectangle2.Area(), rectangle3.Area())
	fmt.Println(rectangle1.IsSquare(), rectangle2.IsSquare(), rectangle3.IsSquare())
	rectangle1.DoubleSize()
	rectangle2.DoubleSize()
	rectangle3.DoubleSize()
	seconds1.Minutes()
	seconds2.Minutes()
	persons1.FullCard()

}
func (r Rectangle) Area() int {
	var a = r.Weight * r.Height
	return a
}

func (r Rectangle) IsSquare() bool {
	return r.Height == r.Weight
}

func (r Rectangle) DoubleSize() {
	var doubleHeight = r.Height * 2
	var doubleWeight = r.Weight * 2
	fmt.Printf("Размеры удвоенного прямоугольника: %d на %d\n", doubleWeight, doubleHeight)
}

func (s Seconds) Minutes() {
	var second = float64(s.Seconds)
	var minute float64 = second / 60.00
	fmt.Printf("У нас есть %f минут.\n", minute)
}

func (a Address) FullAddresse() string {
	adress := a.Street + " " + a.City + " " + a.Country
	return adress
}

func (p Persons) FullCard() {
	fmt.Printf("Name: %s\nAddress: %s\n", p.Name, p.Live.FullAddresse())
}
