package main

import (
	"fmt"
)

// manejo de errores con panic y recover

func dividirConPanic(a, b float64) float64 {
	if b == 0 {
		panic("division por cero") //panic es una funcion que detiene la ejecucion del programa y lanza un error
	}
	return a / b
}

func manejarPanic() {
	if r := recover(); r != nil {
		fmt.Println("Ocurrio un error:", r)
	}
}

func main() {
	defer manejarPanic()

	fmt.Println("inicio de la ejecucion")
	resultado := dividirConPanic(5, 0)
	fmt.Println("Resultado de la division", resultado)
	fmt.Println("Fin programa")
}
