package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) >= 3 {
		s := os.Args[len(os.Args)-1]
		os.Args[len(os.Args)-1] = ""
		for i, arg := range os.Args[1:] {
			if arg == s {
				fmt.Println(i)
				return
			}
		}
		fmt.Println("-1")
		return
	}
	fmt.Println("error")
}
