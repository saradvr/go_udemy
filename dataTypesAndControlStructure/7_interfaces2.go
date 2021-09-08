package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func interfaces2() {
	fmt.Println("")
	Anything(2.44)
	Anything("oewiqjw")
	Anything(struct{}{}) // struct vacío de cualquier tipo. Si fuera a definirlo sería struct{Person string} {"hi"}

	mymap := make(map[string]interface{}) // el key del map va a ser siempre un string pero el valor puede ser cualquier cosa
	mymap["name"] = "Sara"
	mymap["age"] = 10
	Anything(mymap)
}
