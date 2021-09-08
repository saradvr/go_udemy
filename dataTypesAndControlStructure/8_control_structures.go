package main

import "fmt"

func constrolStructures() {
	fmt.Println("Hello World")
	f := true
	flag := &f

	if flag == nil { //estar seguros de que la variable está guardada en algún lugar de memoria
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for i := 0; i < 10; i++ { // si sólo pongo el for, es un ciclo infinito
		fmt.Println(i)
	}

	arr := []string{"sara", "del", "valle"}
	for i, value := range arr { // si sólo pongo i se refiere al índice, si no voy a usar el índice
		fmt.Println(i, value)
	}

	mymap := make(map[string]interface{}) // el key del map va a ser siempre un string pero el valor puede ser cualquier cosa
	mymap["name"] = "Sara"
	mymap["age"] = 10

	for i, value := range mymap {
		fmt.Printf("Key: %s and Value: %v \n", i, value)
	}
}
