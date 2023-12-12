package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(sun(1, 2))
	valor, err := sun2(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(valor)
	}
	fmt.Println(sun2(10, 2))
}

func sun(a, b int) int {
	return a + b
}

func sun2(a, b int) (int, error) {
	if a+b > 10 {
		return a + b, errors.New("Valor maior que 10")
	} else {
		return a + b, nil
	}
}
