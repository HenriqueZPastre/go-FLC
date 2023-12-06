package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Hello"
	e float64 = 3.14
	f ID      = 1
)

func main() {
	fmt.Printf("o tipo de f é %T", f)
}
