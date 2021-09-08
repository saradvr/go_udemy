package main

import "fmt"

func arrays() {
	// var arr []int
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr) // [1 2 3 4]

	arr2 := []string{"hi", "my", "name"}
	arr2 = append(arr2, "is Sara")
	fmt.Println(arr2, arr2[3]) // [hi my name is Sara] is Sara
}
