package main

import (
	"fmt"
	"os"
	"sort"
)

/*
Créez un programme qui trie les éléments donnés en argument par ordre ASCII.


Exemples d’utilisation :
$> python exo.py Alfred Momo Gilbert
Alfred Gilbert Momo


$> python exo.py A Z E R T Y
A E R T Y Z

Afficher error et quitter le programme en cas de problèmes d’arguments.
*/

func main() {
	if len(os.Args) >= 3 {
		arr := os.Args[1:]
		sort.Slice(arr, func(i, j int) bool {
			return int(rune(arr[i][0])) < int(rune(arr[j][0]))
		})
		fmt.Println(arr)
		return
	}
	fmt.Println("error")
}
