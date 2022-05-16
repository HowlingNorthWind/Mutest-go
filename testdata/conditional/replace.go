//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 3
	if a > b && b > c || a < b && b < c {
		t := c
		c = a
		a = t
	}
	fmt.Println("done")
}
