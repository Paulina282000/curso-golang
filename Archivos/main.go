//colecciones de datos
//arreglos, slices, maps, structs, etc

//arreglos, son colecciones de datos que tienen un tamaño fijo
//slices, son colecciones de datos que tienen un tamaño variable
//maps, son colecciones de datos que tienen un tamaño variable y se indexan con una clave
//structs, son colecciones de datos que tienen un tamaño variable y se indexan con una clave

//arrays
//una vez definido el tamaño, no se puede cambiar
//sintaxis del array: var nombreArray [tamaño]tipoDeDato

package main

import "fmt"

func main() {
	var nombres [3]string
	nombres[0] = "Juan"
	nombres[1] = "Maria"
	nombres[2] = "Pedro"

	apellidos := [3]string{"Perez", "Gomez", "Lopez"}

	fmt.Println(nombres)
	fmt.Println(apellidos)
	fmt.Println(nombres[0])
	fmt.Println(apellidos[0])

	//slices
	//son colecciones de datos que tienen un tamaño variable, son dinamicos y mas usados que los arrays
	//sintaxis del slice: var nombreSlice []tipoDeDato

	var frutas []string
	frutas = append(frutas, "Manzana", "pera", "Naranja")

	frutas_tropicales := []string{"Fresa", "Uva", "Mango"}

	fmt.Println(frutas)
	fmt.Println(frutas_tropicales)

	frutas = append(frutas, frutas_tropicales...)

	fmt.Println(frutas)

	//maps
	//son colecciones de datos que tienen un tamaño variable y se indexan con una clave
	//los mapas en go son colecciones desordenadas de pares clave valor
	//vamoas a crear una coleccion que no va a estar ordenada, y se van a crear dos tipos de datos
	//la clave y el valor, similar a los diccionarios de python.
	//permiten almacenar datos en forma de clave valor, y no como los slices que son indexados numericamente
	//sintaxis del map: var nombreMap map[tipoDeDato]tipoDeDato

	var edades map[string]int
	edades = make(map[string]int)

	edades["Juan"] = 20
	edades["Maria"] = 25
	edades["Pedro"] = 30

	fmt.Println(edades)
	fmt.Println(edades["Juan"]) // va la clave y da como resultado 20

	edades_new := map[string]int{
		"Juan":  20,
		"Maria": 25,
		"Pedro": 30,
	}

	fmt.Println(edades_new)
	fmt.Println(edades_new["Juan"])

	//uso de rango
	//Range, se utiliza en go para recorrer los elementos de un slice, array o map
	//sintaxis: for indice, valor := range coleccion

	//devuelve dos valores, el indice y el valor
	//si no queremos el indice, podemos usar el guion bajo

	numeros := []int{10, 20, 30, 40, 50}

	for i, num := range numeros {
		fmt.Println(i, num)
	}
	//SIN INDICE
	for _, num := range numeros {
		fmt.Println(num)
	}

	//Range con mapas
	capitales := map[string]string{
		"Argentina": "Buenos Aires",
		"Brasil":    "Brasilia",
		"Chile":     "Santiago",
	}
	for pais, capital := range capitales {
		fmt.Println(pais, capital)
	}

}
