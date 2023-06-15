package main

import (
	"fmt"
	"strings"
)

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
