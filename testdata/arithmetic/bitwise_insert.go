// +build example-main

package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 0
	c = a & b
	c = a | b
	c = a ^ b
	c = a &^ b
	fmt.Println(c)
}
