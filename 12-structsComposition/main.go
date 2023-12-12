package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

func main() {

	wesley := Cliente{
		Nome:  "Wesley",
		Idade: 30,
		Ativo: true,
	}
	wesley.Cidade = "Rua dos Bobos"
	wesley.Endereco.Cidade = "Rua dos Bobos"

	wesley.Ativo = false
	fmt.Println(wesley)

}
