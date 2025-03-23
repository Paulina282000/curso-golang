package main //nombre del packete, define que el achivo pertenece a main
//en goi si o si tenes que pertenecer a un paquete

import (
	"fmt" //importa el paquete fmt, que es el paquete de entrada y salida de datos
) //para usar println

func main() { //la funcion se encarga de ejecutar el programa
	fmt.Println("Hello, World desde Go!")

	//variables
	//declaracion explicita
	var nombre string
	nombre = "Juan"
	fmt.Println(nombre)

	//declaracion y asignacion simulatanea

	var apellido string = "Perez"
	fmt.Println(apellido)

	//declaracion implicita

	edad := 25
	fmt.Println(edad)

	/*
		constantes son espacion de menmoria donde se guarda un valor que no puede cambiar una
		vez que le asignamos un valor, no se puede cambiar

		const nombre string = "Juan"
		por ejemplo un numero de identificacion
	*/

	const id int = 1234567890
	fmt.Println("El numero de identificacion es:", id)

	const pi float64 = 3.14159
	fmt.Println("El valor de pi es:", pi)

	//tipos de datos en GO

	//numeros enteros
	// uint8, uint16, uint32, uint64, int8, int16,
	// int32, int64

	var entero8 uint8 = 255
	fmt.Println("El valor de entero8 es:", entero8)

	var entero16 uint16 = 65535
	fmt.Println("El valor de entero16 es:", entero16)

	//numero flotantes
	// float32, float64

	var float32 float32 = 2.5
	fmt.Println("El valor de float32 es:", float32)

	var float64 float64 = 2.5
	fmt.Println("El valor de float64 es:", float64)

	//cadena de caracteres
	// string

	var cadena string = "Hola, mundo!"
	fmt.Println("La cadena es:", cadena)

	//booleanos verdaros o falsos
	// bool

	var booleano bool = true
	fmt.Println("El valor de booleano es:", booleano)

	//operadores aritmeticos
	// +, -, *, /, %

	num1 := 10
	num2 := 20

	suma := num1 + num2
	fmt.Println("La suma de", num1, "y", num2, "es:", suma)

	resta := num1 - num2
	fmt.Println("La resta de", num1, "y", num2, "es:", resta)

	multiplicacion := num1 * num2
	fmt.Println("La multiplicacion de", num1, "y", num2, "es:", multiplicacion)

	division := num1 / num2
	fmt.Println("La division de", num1, "y", num2, "es:", division)

	modulo := num1 % num2
	fmt.Println("El modulo de", num1, "y", num2, "es:", modulo)

	//Operdores logicos && and, || or, ! not

	adulto := true
	tienePermiso := false

	puedeEntrar := adulto && tienePermiso
	fmt.Println("No puede entrar al bar:", puedeEntrar)

	puedeSalir := adulto || tienePermiso
	fmt.Println("Puede salir del bar:", puedeSalir)

	noPuedeSalir := !adulto
	fmt.Println("No puede salir del bar:", noPuedeSalir)

	//operadores de comparacion
	// ==, !=, >, <, >=, <=

}

/*
comentario de varias lineas,
no afectan el compilador
*/

/* las variables y las constantes se declaran con var y const
y se utilizan para guardar datos en memoria

*/
