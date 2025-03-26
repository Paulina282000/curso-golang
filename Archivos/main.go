package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// estamos mapearando un json a una estructura de datos
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	fileName := "person.json"

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer file.Close()

	var persona Person

	err = json.NewDecoder(file).Decode(&persona) // si surge un error, se guarda en la variable err

	fmt.Println("Nombre:", persona.Name)
	fmt.Println("Edad:", persona.Age)

}
