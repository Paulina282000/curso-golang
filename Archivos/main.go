package main

import "fmt"

type Empleado struct {
	Nombre   string
	Apellido string
	Edad     int
	Email    string
}

func (e Empleado) MostrarDetalles() {
	fmt.Printf("Nombre: %s\n", e.Nombre)
	fmt.Printf("Apellido: %s\n", e.Apellido)
	fmt.Printf("Edad: %d\n", e.Edad)
	fmt.Printf("Email: %s\n", e.Email)
}

type Gerente struct {
	Empleado
	Departamento string
}

func (g Gerente) MostrarDetallesGerente() {
	g.MostrarDetalles()
	fmt.Printf("Departamento: %s\n", g.Departamento)
}

func main() {
	emp := Empleado{
		Nombre:   "Juan",
		Apellido: "Perez",
		Edad:     30,
		Email:    "juan.perez@example.com",
	}

	ger := Gerente{
		Empleado: Empleado{
			Nombre:   "Ana",
			Apellido: "Gomez",
			Edad:     53,
			Email:    "ANA.GOMEZ@example.com",
		},
		Departamento: "IT",
	}
	emp.MostrarDetalles()
	ger.MostrarDetallesGerente()
}

/*
	Embedding sintaxis:

	type EstruturaA struct {
		campo1 string
		campo2 int
	}

	type EstruturaB struct {
		EstruturaA
		campo3
	}



*/
