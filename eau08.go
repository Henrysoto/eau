package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
Créez un programme qui met en majuscule la première lettre de chaque mot d’une chaîne de caractères. Les autres lettres devront être en minuscules. Les mots ne sont délimités que par un espace, une tabulation ou un retour à la ligne.


Exemples d’utilisation :
$> python exo.py “bonjour mathilde, comment vas-tu ?!”
Bonjour Mathilde, Comment Vas-tu ?!


$> python exo.py 42
error

Afficher error et quitter le programme en cas de problèmes d’arguments.
*/

func main() {
	if len(os.Args) == 2 {
		if unicode.IsLetter(rune(os.Args[1][0])) {
			str := ""
			for _, c := range strings.Fields(os.Args[1]) {
				str = fmt.Sprintf("%s %s%s", str, strings.ToUpper(string(c[0])), c[1:])
			}
			fmt.Println(str)
			return
		}
	}
	fmt.Println("error")
}
