package main

import "fmt"

func main() {
	lista := map[string]int{
		"aluno1": 5,
		"aluno2": 2,
		"aluno3": 8,
		"aluno4": 5,
	}
	nota_aluno := lista["aluno"] //+ lista["aluno2"] + lista["aluno3"] + lista["aluno4"]

	fmt.Println(nota_aluno / len(lista))
}
