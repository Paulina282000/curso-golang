package main //nombre del packete, define que el achivo pertenece a main
//en goi si o si tenes que pertenecer a un paquete

import (
	"fmt" //importa el paquete fmt, que es el paquete de entrada y salida de datos
) //para usar println

func main2() { //la funcion se encarga de ejecutar el programa
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

	//estructuras de control: permiten controlar el flujo de ejecucion del programa

	//if,else switch --> if si es verdadera, else si es falsa, switch es una estructura condicional que se encarga de ejecutar un bloque de codigo si una condicion es verdadera

	edadEstudiante := 17

	if edadEstudiante >= 18 {
		fmt.Println("El estudiante es mayor de edad")
	} else {
		fmt.Println("El estudiante es menor de edad")
	}

	//condicional switch, se usa para simplificar multiples condiciones

	dia := 10
	switch dia {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miercoles")
	case 4:
		fmt.Println("Jueves")

	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sabado")
	case 7:
		fmt.Println("Domingo")
	default:
		fmt.Println("Numero invalido")
	}
	//BUCLE O CICLO FOR, sirve para repetir un bloque de codigo un numero determinado de veces

	//for tradicional

	//for inicializaion,donde empieza; condicion,donde termina;incremente,donde se incrementa{}

	for i := 0; i <= 10; i++ {
		fmt.Println("El valor de i es:", i)
	}

	//for con rango, itera sobre listas,arreglos,mapas,etc

	nombres := []string{"Juan", "Maria", "Pedro", "Ana"}

	for i, nombre := range nombres {
		fmt.Println("El valor de i es:", i, "y el valor de nombre es:", nombre)
	}

	//funciones, son bloques de codigo que se pueden reutilizar, modularizar, y encapsular
	//pueden tener parametros y devolver valores
	//Sintaxis:
	//fun nombreFuncion(parametros) tiporetorno{
	//	codigo }

	//llamo la funcnion sumar desde el main

	resultado, resultado2 := sumarYRestar(10, 20)
	fmt.Println("El resultado de la suma es:", resultado)
	fmt.Println("El resultado de la resta es:", resultado2)
}

// las funcines deben estar fuera del main par que funcionen
func sumarYRestar(num1 int, num2 int) (int, int) {
	return num1 + num2, num1 - num2
}
