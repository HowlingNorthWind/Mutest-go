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
	fmt.Println("package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\ti := 1\n\n\tfor i != 4 {\n\t\tswitch {\n\t\tcase i == 1:\n\t\t\tfmt.Println(i)\n\t\tcase i == 2:\n\t\t\tfmt.Println(i * 2)\n\t\tcase i == 3:\n\t\t\tfmt.Println(i * 3)\n\t\t\tfallthrough\n\t\t\tfallthrough\n\t\tdefault:\n\t\t\tfmt.Println(i * 4)\n\t\t}\n\n\t\ti++\n\t}\n}\n")
}
