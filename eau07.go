package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
Créez un programme qui met en majuscule une lettre sur deux d’une chaîne de caractères. Seule les lettres (A-Z, a-z) sont prises en compte.


Exemples d’utilisation :
$> python exo.py “Hello world !”
HeLlO wOrLd !


$> python exo.py 42
error

Afficher error et quitter le programme en cas de problèmes d’arguments.
*/

func main() {
	if len(os.Args) == 2 {
		if len(strings.TrimSpace(os.Args[1])) > 0 && !unicode.IsDigit(rune(os.Args[1][0])) {
			str := ""
			for i, x := range strings.Split(os.Args[1], "") {
				if i%2 == 0 {
					str = fmt.Sprintf("%s%s", str, strings.ToUpper(x))
				} else {
					str = fmt.Sprintf("%s%s", str, x)
				}
			}
			fmt.Println(str)
		} else {
			fmt.Println("erreur.")
		}
	} else {
		fmt.Println("erreur.")
	}
}
