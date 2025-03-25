package main

import "fmt"

//uso de intefaces en go

type Hablador interface {
	Hablar() string
}

type Persona struct {
	Nombre string
	Edad   int
}

func (p *Persona) Hablar() string { //aca vemos como se implementa la interface Hablador, para que la struct Persona pueda usar el metodo Hablar
	return "Hola, mi nombre es " + p.Nombre
}

func main() {
	persona := Persona{Nombre: "Juan", Edad: 20}
	fmt.Println(persona.Hablar())
}
