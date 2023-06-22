package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Créez un programme qui affiche toutes les valeurs comprises entre deux nombres dans l’ordre croissant. Min inclus, max exclus.


Exemples d’utilisation :
$> python exo.py 20 25
20 21 22 23 24


$> python exo.py 25 20
20 21 22 23 24

$> python exo.py hello
error

Afficher error et quitter le programme en cas de problèmes d’arguments.
*/

func main() {
	if len(os.Args) == 3 {
		min, _ := strconv.ParseInt(os.Args[1], 10, 32)
		max, _ := strconv.ParseInt(os.Args[2], 10, 32)
		if max < min {
			tmp := min
			min = max
			max = tmp
		}
		for ; min < max; min++ {
			fmt.Printf("%d ", min)
		}
		fmt.Println()
		return
	}
	fmt.Println("error")
}
