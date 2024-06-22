package main

import "fmt"

func main() {
	// declaracion de constantes:
	const pi float64 = 3.14
	const pi2 = 3.16

	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	// declaracion de variables enteras
	// primera forma
	base := 12
	// segunda forma
	var altura int = 14
	// tercera forma
	var area int
	//Go no compila si las variables no son usadas
	fmt.Println(base, altura, area)

	//Zero values
	//Go asigna vaalores a variables vacías
	var a int     // 0
	var b float64 // 0
	var c string  // string vacio
	var d bool    // false

	fmt.Println(a, b, c, d)

	//Ejercicio
	//Calcular el áera del cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

}
