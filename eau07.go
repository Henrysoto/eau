package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

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
