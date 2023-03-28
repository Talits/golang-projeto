package main

import "fmt"

func soma(a int, b int) int {
	return a + b
}

func imprimir(resultado int) {
	fmt.Println(resultado)
}

func main() {
	resultado := soma(1, 2)
	imprimir(resultado)
}
