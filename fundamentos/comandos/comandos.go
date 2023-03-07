package main

import (
	"fmt"
)

func main() {
	const PI float64 = 3.1415
	const (
		a = 1
		b = 3
	)
	g, h, i := 1, false, "teste"
	area := 3.12
	fmt.Println("Testando", area)
	fmt.Print(a, b)
	fmt.Println(g, h, i)
}
