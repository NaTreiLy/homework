package main

import "fmt"

func main() {

	garden := map[string]int{
		"lemon":  3,
		"cherry": 6,
		"apple":  2,
	}
	fmt.Println(garden)

	garden["cherry"] = 5

	fmt.Println(garden)

	fmt.Printf("В нашем саду %d вишен", garden["cherry"])

	garden["oak"] = 7
	garden["birch"] = 3
	fmt.Println(garden)
	delete(garden, "oak")
	fmt.Println(garden)

	if tree, ok := garden["birch"]; ok {
		fmt.Printf("Березы есть в нашем саду, их %d", tree)
	} else {
		fmt.Println("Березы стоит посадить")
	}

	for key, value := range garden {
		fmt.Printf("В нашем саду есть: %s - %d", key, value)
	}
}
