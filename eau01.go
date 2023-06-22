package main

import (
	"fmt"
	"strings"
)

/*
Créez un programme qui affiche toutes les différentes combinaisons possibles de trois chiffres dans l’ordre croissant, dans l’ordre croissant. La répétition est volontaire.


Exemples d’utilisation :
$> python exo.py
012, 013, 014, 015, 016, 017, 018, 019, 023, 024, ... , 789
$>

987 n’est pas là parce que 789 est présent.

020 n’est pas là parce que 0 apparaît deux fois.

021 n’est pas là parce que 012 est présent.

000 n’est pas là parce que cette combinaison ne comporte pas exclusivement des chiffres différents les uns des autres.
*/

func contains(arr []string, str string) bool {
	for _, i := range arr {
		if i == str {
			return true
		}
		for j := 0; j <= 2; j++ {
			for k := 0; k <= 2; k++ {
				for l := 0; l <= 2; l++ {
					x := fmt.Sprintf("%c%c%c", i[j], i[k], i[l])
					if x == str {
						return true
					}
				}
			}
		}
	}
	return false
}

func main() {
	var s string
	var t []string
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			if i != j {
				for k := 0; k <= 9; k++ {
					if i != k && j != k {
						s = fmt.Sprintf("%d%d%d", i, j, k)
						if len(s) == 3 && !contains(t, s) {
							t = append(t, s)
						}
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", strings.Join(t, ", "))
}
