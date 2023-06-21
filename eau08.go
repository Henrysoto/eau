package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		str := ""
		for _, c := range strings.Fields(os.Args[1]) {
			str = fmt.Sprintf("%s %s%s", str, strings.ToUpper(string(c[0])), c[1:])
		}
		fmt.Println(str)
	} else {
		fmt.Println("erreur.")
	}
}
