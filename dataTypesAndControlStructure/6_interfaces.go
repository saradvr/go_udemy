package main

import "fmt"

type Car2 interface {
	Drive2()
	Stop()
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive2() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive2() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func interfaces() {
	l := Lambo{"Gallardo"}
	c := Chevy{"C369"}
	l.Drive2()
	c.Drive2()
}
