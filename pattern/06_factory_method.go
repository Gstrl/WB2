package main

import (
	"errors"
	"fmt"
)

/*
Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Применимость паттерна "фабричный метод":
1. Заранее не известны точные типы и зависимости объектов
2. Возможность расширения внутренних компонентов
3. Экономия системных ресурсов, повторно используя существующие объекты

Плюсы паттерна "фабричный метод":
1. позволяет сделать код создания объектов более универсальным, не привязываясь к конкретным классам
2. Принцип единой ответственности: Можно переместить код создания продукта в одно место программы, что упростит поддержку кода.
3. Принцип открытости/закрытости: Можно вводить в программу новые виды продуктов, не нарушая существующий код.

Минусы паттерна "фабричный метод":
1. Необходимость создавать наследника Creator для каждого нового типа продукта
2. Код может стать более сложным: В небольших программах добавление фабричных методов может привести к избыточности кода и усложнению его структуры
*/

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
