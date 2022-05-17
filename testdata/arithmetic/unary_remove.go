// +build example-main

package main

import "fmt"

func main() {
	i := 100

	i = +i
	i = -i
	i = ^i
	z := false
	z = !z

	fmt.Println(i)
}
