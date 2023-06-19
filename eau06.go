package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func strContains(str string, substr string) bool {
	if len(str) > 0 && len(substr) > 0 {
		w := strings.Split(substr, "")
		x := strings.Split(str, "")
		t := 0
		vi := 1
		for i := 1; i <= len(w); i++ {
			for j := vi; j <= len(x); j++ {
				if w[i-1] == x[j-1] {
					vi = j + 1
					t++
					//fmt.Printf("\tx='%s' :: w='%s' :: \033[0;32mvalid=%d\033[0m :: index_x=%d :: index_w=%d\n", x[j-1], w[i-1], t, j, i)
					break
				} else {
					t = 0
					vi = 1
					//fmt.Printf("\tx='%s' :: w='%s' :: valid=%d :: index_x=%d :: index_w=%d\n", x[j-1], w[i-1], t, j, i)
					i = 1
					if j == len(x) {
						i = len(w) + 1
						break
					}
				}
			}
		}
		return t >= len(w)
	}
	return false
}

func main() {
	if len(os.Args) == 3 {
		if unicode.IsDigit(rune(os.Args[1][0])) {
			fmt.Println("error")
			return
		}
		if strContains(os.Args[1], os.Args[2]) {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else {
		fmt.Println("error")
	}
}
