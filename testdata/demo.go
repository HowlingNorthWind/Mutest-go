package main

import (
	"fmt"
	"mutesting/testdata/testRec/test"
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
	c := t{
		a,
		"",
		0,
		0,
	}
	fmt.Println(c.b)
	test.Rec(a)
	fmt.Print("//go:build ignore\n// +build ignore\n\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n\ti := 100\n\n\ti = +i\n\ti = -i\n\ti = +i\n\tz := false\n\tz = !z\n\n\tfmt.Println(i)\n}\n")
}
