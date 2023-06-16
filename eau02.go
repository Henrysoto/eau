package main

import "fmt"

func main() {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			if i < j {
				fmt.Printf("%02d %02d\t", i, j)
			}
		}
	}
	fmt.Println()
}
