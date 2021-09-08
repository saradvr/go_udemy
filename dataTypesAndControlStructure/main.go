package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // va al siguiente ciclo
		}
		fmt.Println(i)
	}

	f := true
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			f = false
			break // rompe el ciclo
		}
		fmt.Println(i)
	}
	fmt.Println(f)

	day := "thr"
	switch day {
	case "fri":
		fmt.Println("TGIF")
	case "mon", "tue", "wed":
		fmt.Println("boring")
	default:
		fmt.Println("default")
	}
}
