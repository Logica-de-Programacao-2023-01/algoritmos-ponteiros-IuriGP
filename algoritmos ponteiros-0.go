package main

import "fmt"

func main() {
	var x bool
	var pontr *bool = &x

	x = false

	if x == true {
		x = false
		printValue(pontr)
	} else {
		x = true
		printValue(&x)
	}

}
func printValue(ptr *bool) {
	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("Pointer is nil")
	}
}

type info struct {
	nome             string
	preço            float64
	quantidadeStoque int
}

func Novopreço() {
	
}
