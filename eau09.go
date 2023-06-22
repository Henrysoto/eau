package main

import (
	"fmt"
	"os"
	"regexp"
)

/*
Créez un programme qui détermine si une chaîne de caractères ne contient que des chiffres.


Exemples d’utilisation :
$> python exo.py “4445353”
true


$> python exo.py 42
true

$> python exo.py “Bonjour 36”
false

Afficher error et quitter le programme en cas de problèmes d’arguments.
*/

func main() {
	if len(os.Args) == 2 {
		m := regexp.MustCompile("[[:alpha:]]")
		if ok := m.Match([]byte(os.Args[1])); !ok {
			fmt.Println(!ok)
		} else {
			fmt.Println("false")
		}
		return
	}
	fmt.Println("error")
}
