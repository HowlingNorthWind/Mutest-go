// +build example-main

package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Println(a & ^b)
	fmt.Println(^a | b)
}
