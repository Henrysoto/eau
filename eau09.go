package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) == 2 {
		m := regexp.MustCompile("[[:alpha:]]")
		if ok := m.Match([]byte(os.Args[1])); !ok {
			fmt.Println(!ok)
		} else {
			fmt.Println("false")
		}
		return
	}
	fmt.Println("error")
}
