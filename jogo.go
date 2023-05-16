package main

import (
	"fmt"
)

func main() {
	var resp int
	tentativas := 0
	num := 50

	for {
		fmt.Scan(&resp)

		if resp < num {
			fmt.Println("O num é maior que", resp)
			fmt.Println("tentativas: "), tentativas++

		} else if resp > num {
			fmt.Println("O numero é menor que ", resp)
			tentativas++

		} else if resp == num {
			fmt.Println("certa resposta!")
			tentativas++
			break
		}
		fmt.Println(tentativas)
	}
}
