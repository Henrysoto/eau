package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Créez un programme qui affiche le N-ème élément de la célèbre suite de Fibonacci. (0, 1, 1, 2) étant le début de la suite et le premier élément étant à l’index 0.


Exemples d’utilisation :
$> python exo.py 3
2
$>

Afficher -1 si le paramètre est négatif ou mauvais.
*/

func main() {
	if len(os.Args) == 2 {
		fibonacci := []int64{0, 1}
		arg, _ := strconv.ParseInt(os.Args[1], 10, 64)
		if arg >= 0 {
			for i := int64(2); i <= arg; i++ {
				fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
			}
			fmt.Println(fibonacci[arg])
		} else {
			fmt.Println("-1")
		}
	} else {
		fmt.Println("-1")
	}
}
