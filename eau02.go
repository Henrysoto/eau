package main

import "fmt"

/*
Créez un programme qui affiche toutes les différentes combinaisons de deux nombre entre 00 et 99 dans l’ordre croissant.


Exemples d’utilisation :
$> python exo.py
00 01, 00 02, 00 03, 00 04, ... , 00 99, 01 02, ... , 97 99, 98 99
$>
*/

func main() {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if i < j {
				fmt.Printf("%02d %02d\t", i, j)
			}
		}
	}
	fmt.Println()
}
