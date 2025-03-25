// Go: Ejemplo de interfaces y polimorfismo
package main

import "fmt"

//interface vacia:
//interface{}: tipo vacio, representa cualquier valor
//util para:
//- funciones que pueden recibir cualquier tipo de dato
//- estructuras de datos heterogeneas (ej: JSON)
//- pruebas unitarias (ej: mock objects)

//sintaxis interface vacia:
//interface{}

// Shape define un comportamiento común para cualquier figura
type Shape interface {
	Area() float64
}

// Circle representa un círculo con su radio
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.1416 * c.radius * c.radius
}

// Rectangle representa un rectángulo con ancho y alto
type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// printArea acepta cualquier Shape e imprime su área
func printArea(s Shape) {
	fmt.Println("Área:", s.Area())
}

func main() {
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 3, height: 4}

	printArea(circle)    // Polimorfismo: Circle implementa Shape
	printArea(rectangle) // Polimorfismo: Rectangle implementa Shape
}
