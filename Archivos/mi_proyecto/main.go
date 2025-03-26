package main

import (
	"fmt"
	"mi_proyecto/matematicas"
)

func main() {
	suma := matematicas.Suma(10, 5)
	resta := matematicas.Resta(10, 5)

	fmt.Println("Suma:", suma)
	fmt.Println("Resta:", resta)
}
