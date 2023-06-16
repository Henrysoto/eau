package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 2 {
		for i := len(os.Args) - 1; i > 0; i-- {
			fmt.Printf("%s ", os.Args[i])
		}
		fmt.Println()
	} else {
		fmt.Println("erreur, pas assez d'arguments.")
	}
}
