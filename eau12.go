package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
Créez un programme qui affiche la différence minimum absolue entre deux éléments d’un tableau constitué uniquement de nombres. Nombres négatifs acceptés.


Exemples d’utilisation :
$> python exo.py 5 1 19 21
2


$> python exo.py 20 5 1 19 21
1

$> python exo.py -8 -6 4
2

Afficher error et quitter le programme en cas de problèmes d’arguments.
*/

func main() {
	if len(os.Args) >= 3 {
		var t []int
		for _, n := range os.Args[1:] {
			x, _ := strconv.ParseInt(n, 10, 32)
			t = append(t, int(x))
		}
		var tmp []int
		for i := 0; i < len(t); i++ {
			for j := 0; j < len(t); j++ {
				if t[i] == t[j] {
					continue
				}
				if t[i]-t[j] >= 0 {
					tmp = append(tmp, t[i]-t[j])
				}
			}
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		fmt.Println(tmp[0])
		return
	}
	fmt.Println("error")
}
