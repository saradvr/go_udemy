package main

import "fmt"

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

func swap2(m1, m2 int) { // Sin el pointer no cambia el valor de la posición en memoria, sólo cambia dentro de la función pero no lo que está afuera
	var temp int
	temp = m2
	m2 = m1
	m1 = temp
	fmt.Println("inside swap2", m1, m2) // Intercambia los valores sólo dentro de la función
}

func pointers() {
	m1 := 2000
	ptr := &m1
	fmt.Println(ptr)  // Memory address 0xc000014068
	fmt.Println(*ptr) // Value of the pointer 2000

	m2, m3 := 2, 3
	fmt.Println(m2, m3)   // 2 3
	fmt.Println(&m2, &m3) // 0xc0000140a8 0xc0000140b0
	swap(&m2, &m3)
	fmt.Println(m2, m3)   // 3 2
	fmt.Println(&m2, &m3) // 0xc0000140a8 0xc0000140b0

	swap2(m2, m3)
	fmt.Println("outside swap2", m2, m3) // Mantiene los valores que tenían porque no cambió el valor de la posición en memoria
}
