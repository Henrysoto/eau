package main

import (
	"fmt"
	"os"
	"sort"
)

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
