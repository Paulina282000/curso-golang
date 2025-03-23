package main

import "fmt"

//##Punteros y funciones

func modificarValor(ptr *int) {
	*ptr = 50
}

func main() {
	//declaracion de variables

	x := 10

	fmt.Println("Valor de x antes de la funcion:", x)

	//pasar la direccion de memoria de x a la funcion
	modificarValor(&x)

	fmt.Println("Valor de x despues de la funcion:", x)

}
