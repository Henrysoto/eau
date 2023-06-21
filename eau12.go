package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	if len(os.Args) >= 3 {
		var t []int
		for _, n := range os.Args[1:] {
			x, _ := strconv.ParseInt(n, 10, 32)
			t = append(t, int(x))
		}
		var tmp []int
		for i := 0; i < len(t); i++ {
			for j := 0; j < len(t); j++ {
				if t[i] == t[j] {
					continue
				}
				if t[i]-t[j] >= 0 {
					tmp = append(tmp, t[i]-t[j])
				}
			}
		}
		sort.Slice(tmp, func(i, j int) bool {
			return tmp[i] < tmp[j]
		})
		fmt.Println(tmp[0])
		return
	}
	fmt.Println("error")
}
