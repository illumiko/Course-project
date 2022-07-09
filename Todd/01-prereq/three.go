package main

import "fmt"

type introduce interface {
	introduction() string
}
type person struct {
	name string
	age  int
}
type agent struct {
	person person
	id     int
}

func (p *person) introduction() (spoken string) {
	fmt.Printf("Hello Im %s\n", p.name)
	return
}

func (a *agent) introduction() (spoken string) {
	fmt.Printf("Hello Im agent %v\n", a.id)
	return
}
func able_to_introduce(r introduce) (x string) {
	fmt.Printf("This is of type %T \n", r)
	return
}

func main() {
	p1 := person{"Smith", 20}
	a1 := agent{p1, 2205}
	a1.introduction()
	able_to_introduce(&p1)
}
