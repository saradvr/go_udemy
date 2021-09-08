package main

import "fmt"

// structure (data encapsulation)
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() { // es un método del objeto que le pongo al inicio, se llama con el c.Print()
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving...", c.Name)
}

func (c Car) GetName() string { // se debe poner al final el tipo de dato que retorna
	return c.Name
}

func structures() {
	// var c1 Car
	c := Car{"chevy", 1, 2}
	fmt.Println(c) // {chevy 1 2}

	c1 := Car{
		Name:    "Chevy",
		Age:     2,
		ModelNo: 2020,
	}
	fmt.Println(c1)           // {Chevy 2 2020}
	c1.Print()                // {Chevy 2 2020}
	c1.Drive()                // Driving... Chevy
	fmt.Println(c1.GetName()) // Chevy
}
