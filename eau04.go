package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 2 {
		fibonacci := []int64{0, 1}
		arg, _ := strconv.ParseInt(os.Args[1], 10, 64)
		if arg >= 0 {
			for i := int64(2); i <= arg; i++ {
				fibonacci = append(fibonacci, fibonacci[i-1]+fibonacci[i-2])
			}
			fmt.Println(fibonacci[arg])
		} else {
			fmt.Println("-1")
		}
	} else {
		fmt.Println("-1")
	}
}
