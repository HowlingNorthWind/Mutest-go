package main

import (
	"fmt"
)

func Inc(a int) int {
	return a + 1
}

type t struct {
	b int
	s string
	d int
	f float64
}

// main方法
func main() {
	a := 1
	a = Inc(a)
	Inc(a)
	switch a + 2 {
	case 3:
		fmt.Println("a + 2 = 3")
		fallthrough
	case 4:
		fmt.Println("a + 2 == 4")
	}
}
