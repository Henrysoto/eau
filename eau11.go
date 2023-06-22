package main

import (
	"fmt"
	"os"
)

/*
Créez un programme qui affiche le premier index d’un élément recherché dans un tableau. Le tableau est constitué de tous les arguments sauf le dernier. L’élément recherché est le dernier argument. Afficher -1 si l’élément n’est pas trouvé.


Exemples d’utilisation :
$> python exo.py Ceci est le monde qui contient Charlie un homme sympa Charlie
6


$> python exo.py test test test
0

$> python exo.py test boom
-1

Afficher error et quitter le programme en cas de problèmes d’arguet.
*/

func main() {
	if len(os.Args) >= 3 {
		s := os.Args[len(os.Args)-1]
		os.Args[len(os.Args)-1] = ""
		for i, arg := range os.Args[1:] {
			if arg == s {
				fmt.Println(i)
				return
			}
		}
		fmt.Println("-1")
		return
	}
	fmt.Println("error")
}
