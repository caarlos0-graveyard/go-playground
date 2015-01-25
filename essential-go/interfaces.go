package main

import "fmt"

type Animal interface {
	Pet()
	Name() string
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("woof woof")
}

func (d Dog) Name() string {
	return d.name
}

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Println("prrr")
}

func (c Cat) Name() string {
	return c.name
}

type Chinchila struct {
	name string
}

func (c Chinchila) Pet() {
	fmt.Println("rrrrr")
}

func (c Chinchila) Name() string {
	return c.name
}

func Compliment(a Animal) {
	fmt.Println("Great Job", a.Name())
	a.Pet()
}

func Treat(a Animal) {
	fmt.Println("Giving food to", a.Name())
	a.Pet()
}

func main() {
	c := Cat{"Random cat"}
	d := Dog{"Dog"}
	z := Chinchila{"Zakk"}
	Compliment(c)
	Compliment(d)
	Compliment(z)
	Treat(z)
}
