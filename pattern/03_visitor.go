package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Применимость паттерна "посетитель":
1. Необходимо выполнить некоторую (одну и ту же) операцию для ряда объектов, без добавления этой операции в сам класс каждого объекта.
2. Определить новую операцию над объектами без изменения их структуры.

Плюсы паттерна "посетитель":
1. Принцип открытости/закрытости. Можно ввести новое поведение, которое может работать с объектами различных классов, не изменяя эти классы.
2. Принцип единой ответственности. Можно перенести несколько версий одного и того же поведения в один класс.

Минусы паттерна "посетитель":
1. Если структура элементов часто меняется, приходится изменять все посетителей
2. Усложнение кода из-за введения дополнительных классов и интерфейсов
*/

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
