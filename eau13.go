package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

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
