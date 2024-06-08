package main

import "fmt"

//Шаблон посетителя — это шаблон проектирования,
//который позволяет добавлять новые операции к коллекции объектов без (и это важно) изменения самих объектов.

type Visitor interface {
	VisitConcreteElementA()
	VisitConcreteElementB()
}

type ConcreteVisitor struct{}

func (cv ConcreteVisitor) VisitConcreteElementA() {
	fmt.Println("visit element A")
}

func (cv ConcreteVisitor) VisitConcreteElementB() {
	fmt.Println("visit element B")
}

type Element interface {
	Accept(v Visitor)
}

type ConcreteElementA struct{}

type ConcreteElementB struct{}

func (el ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA()
}

func (el ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB()
}

func main() {
	v := Visitor(ConcreteVisitor{})
	//ConcreteElementA
	el := Element(ConcreteElementA{})
	el.Accept(v)
	//ConcreteElementB
	el = Element(ConcreteElementB{})
	el.Accept(v)
}
