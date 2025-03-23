package main

import "fmt"

//CRECION DE FUNCIONES RECURSIVAS

func suma(n int) int {
	if n == 0 {
		return 0
	}
	return n + suma(n-1)
}

func main() {

	//recursion  es una tecnica de la programacion
	//donde una funcion se llama a si misma de forma directa o indirecta,
	//para resolver un problema, subdividiendolo en subproblemas mas pequeños

	//en GO se implementa definiendo una funcion que tiene una condicion base para
	//detener la recursion y una llamada recursiva para resolver el problema

	//codicion base: hasta cuando se puede llamar para no entrar en bucle infinito
	//llamada recursiva: se llama a la funcion dentro de si misma, con argumentos que cada ves que
	//se llamen van a se modificados

	//sumar un numero de 1 a n

	numero := 5
	resultado := suma(numero)
	fmt.Println("La suma de los numeros de 1 a", numero, "es:", resultado)

}
