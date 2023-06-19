package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func prime(x float64) bool {
	if int(x) < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(x)); i++ {
		if int(x)%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) == 2 {
		if int(os.Args[1][0]) == int('-') {
			fmt.Println("-1")
			return
		}
		a, _ := strconv.ParseFloat(os.Args[1], 10)
		for i := a + 1.0; ; i++ {
			if prime(i) {
				fmt.Println(int(i))
				return
			}
		}
	} else {
		fmt.Println("-1")
	}
}
