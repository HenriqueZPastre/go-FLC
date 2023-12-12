package main

import "fmt"

func main() {

	/// Hashtables Maps

	salarios := map[string]int{
		"Wesley": 11325,
	}
	fmt.Println(salarios["Wesley"])
	delete(salarios, "Wesley")
	fmt.Println(salarios["Wesley"])
	salarios["Henrique"] = 5000
	fmt.Println(salarios["Henrique"])

	//sal := make(map[string]int)
	//sal2 := map[string]int{}

	for nome, salario := range salarios {
		fmt.Printf("O salário de %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("O salário é %d\n", salario)
	}
}
