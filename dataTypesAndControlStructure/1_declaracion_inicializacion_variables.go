package main

import "fmt"

func declaracionInicializacion() {

	var m1 int
	m1 = 2

	fmt.Println(m1)

	var (
		m2 = 2
		m3 = 3
	) // con el igual el tipo queda implícito
	fmt.Println(m2 + m3)

	var m4 int32 // si no se pone un valor, se inicializa con 0
	var m5 int64
	//fmt.Println(m4 + m5) //No permite sumar diferentes tipos
	fmt.Println(int64(m4) + m5)

	m6 := 6
	m7 := 7 // es lo mismo que var m7 int m7 = 7
	fmt.Println(m6 + m7)

}
