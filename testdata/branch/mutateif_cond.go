package main

import (
	"fmt"
)

var m = map[int]bool{
	3: true,
}

func main() {
	i := 1

	for i != 4 {
		if _, ok := m[i]; ok {
			fmt.Println(i * 3)
		}
		if i == 1 {
			fmt.Println(i)
		} else if i == 2 {
			fmt.Println(i * 2)
		} else {
			fmt.Println(i * 4)
		}

		i++
	}
}
