package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 {
		min, _ := strconv.ParseInt(os.Args[1], 10, 32)
		max, _ := strconv.ParseInt(os.Args[2], 10, 32)
		if max < min {
			tmp := min
			min = max
			max = tmp
		}
		for ; min < max; min++ {
			fmt.Printf("%d ", min)
		}
		fmt.Println()
		return
	}
	fmt.Println("error")
}
