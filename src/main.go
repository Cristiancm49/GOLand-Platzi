package main

import "fmt"

func main() {

	const pi float64 = 3.14
	const pi2 = 3.1514
	
	fmt.Println("pi: ", pi)
	fmt.Println("pi2: ", pi2)

	// Declaracion de variables

	base := 12
	var altura int =14
	var area int

	fmt.Println(base, altura, area)

	var a int
	var b string
	var c float32
	var d bool

	fmt.Println(a, b, c, d)

	// Area cuadrado

	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("area cuadrado: ", areaCuadrado)
}