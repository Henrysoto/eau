package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
Créez un programme qui trie une liste de nombres. Votre programme devra implémenter l’algorithme du tri à bulle.

Vous utiliserez une fonction de cette forme (selon votre langage) :
my_bubble_sort(array) {
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

func bsort(list []int, n int) {
	for true {
		for i := 1; i <= n-1; i++ {
			if list[i-1] > list[i] {
				tmp := list[i-1]
				list[i-1] = list[i]
				list[i] = tmp
			}
		}
		n = n - 1
		if n == -1 {
			break
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
		fmt.Println(arr, "not sorted")
		bsort(arr, len(arr))
		fmt.Println(arr, "sorted")
		return
	}
	fmt.Println("error")
}
