package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
Créez un programme qui trie une liste de nombres. Votre programme devra implémenter l’algorithme du tri par sélection.

Vous utiliserez une fonction de cette forme (selon votre langage) :
my_select_sort(array) {
	# votre algorithme
	return (new_array)
}


Exemples d’utilisation :
$> python exo.py 6 5 4 3 2 1
1 2 3 4 5 6


$> python exo.py test test test
error

Afficher error et quitter le programme en cas de problèmes d’arguments.


Wikipedia vous présentera une belle description de cet algorithme de tri.
*/

func selSort(list []int, n int) {
	for i := 0; i <= n-2; i++ {
		min := i
		for j := i + 1; j <= n-1; j++ {
			if list[j] < list[min] {
				min = j
			}
		}
		if min != i {
			tmp := list[min]
			list[min] = list[i]
			list[i] = tmp
		}
	}
}

func main() {
	if len(os.Args) >= 3 {
		arr := make([]int, 0, len(os.Args[1:]))
		for _, x := range os.Args[1:] {
			if unicode.IsLetter(rune(x[0])) {
				fmt.Println("error")
				return
			}
			y, _ := strconv.ParseInt(x, 10, 32)
			arr = append(arr, int(y))
		}
		selSort(arr, len(arr))
		fmt.Println(arr)
		return
	}
	fmt.Println("error")
}
