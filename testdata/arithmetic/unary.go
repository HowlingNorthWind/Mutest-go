//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	i := 100

	i = +i
	i = -i

	fmt.Println(i)
}
