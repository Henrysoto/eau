package main

import (
	"fmt"
	"os"
)

/*
Créez un programme qui affiche ses arguments reçus à l’envers.


Exemples d’utilisation :
$> python exo.py “Suis” “Je” “Drôle”
Drôle
Je
Suis


$> python exo.py ha ho
ho
ha

$> python exo.py “Bonjour 36”
Bonjour 36

Afficher error et quitter le programme en cas de problèmes d’arguments.
*/

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
