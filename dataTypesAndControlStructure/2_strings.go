package main

import (
	"fmt"
	"strings"
)

func stringsMain() {

	var m1 string
	m1 = "my name"
	m1 = "Sara"
	fmt.Println(m1) // Sara

	m2 := "my name"
	//m2 := "Sara" // No permite reinicializarse así nuevamente
	m2 = "Sara"
	fmt.Println(m2) // Sara

	m3 := "Hola"
	m4 := "Mundo"
	m5 := "Hola Mundo"
	fmt.Println(strings.Contains(m3, m4))            // false
	fmt.Println(strings.Contains(m5, m4))            // true
	fmt.Println(strings.ReplaceAll(m3, "H", "Hooo")) // Hooola
	fmt.Println(strings.Split(m5, " "))              // los pone en un array [Hola Mundo]

}
