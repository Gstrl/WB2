package main

import (
	"errors"
	"fmt"
)

type Creator interface {
	New(typeProductF string) (ProductF, error)
}

type ProductF interface {
	Print()
}

type ConcreteCreator struct {
}

type ConcreteProductFA struct {
	name  string
	count int
}

func (c ConcreteProductFA) Print() {
	fmt.Printf("ConcreteProductFA: name %s, count %d\n", c.name, c.count)
}

func NewConcreteProductFA() ProductF {
	return ConcreteProductFA{
		name:  "type A designed by builder",
		count: 123,
	}
}

type ConcreteProductFB struct {
	name string
}

func (c ConcreteProductFB) Print() {
	fmt.Printf("ConcreteProductFB: name %s\n", c.name)
}

func NewConcreteProductFB() ProductF {
	return ConcreteProductFB{
		name: "type B designed by builder",
	}
}

func (c ConcreteCreator) New(typeProductF string) (ProductF, error) {
	switch typeProductF {
	default:
		fmt.Printf("%s  не сущесвует\n", typeProductF)
	case "A":
		return NewConcreteProductFA(), nil
	case "B":
		return NewConcreteProductFB(), nil
	}
	return nil, errors.New("не удалось создать ProductF\n")
}

func main() {
	creator := Creator(ConcreteCreator{})

	prA, err := creator.New("A")
	if err != nil {
		fmt.Println(err)
	} else {
		prA.Print()
	}

	prB, err := creator.New("B")
	if err != nil {
		fmt.Println(err)
	} else {
		prB.Print()
	}

	prC, err := creator.New("C")
	if err != nil {
		fmt.Println(err)
	} else {
		prC.Print()
	}
}
