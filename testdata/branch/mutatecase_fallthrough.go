package main

import (
	"fmt"
)

func main() {
	i := 1

	for i != 4 {
		switch {
		case i == 1:
			fmt.Println(i)
		case i == 2:
			fmt.Println(i * 2)
		case i == 3:
			fmt.Println(i * 3)
			fallthrough
		default:
			fmt.Println(i * 4)
		}

		i++
	}
}
