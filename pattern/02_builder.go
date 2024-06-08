package main

import "fmt"

/*
Product (продукт) - Класс, который определяет сложный объект,
который мы пытаемся шаг за шагом сконструировать, используя простые объекты.

Builder (строитель) - абстрактный класс/интерфейс, который определяет все этапы,
необходимые для производства сложного объекта-продукта.
Как правило, здесь объявляются (абстрактно) все этапы (buildPart),
а их реализация относится к классам конкретных строителей (ConcreteBuilder).

ConcreteBuilder (конкретный строитель) - класс-строитель,
который предоставляет фактический код для создания объекта-продукта.
У нас может быть несколько разных ConcreteBuilder-классов, каждый из которых реализует различную разновидность или способ создания объекта-продукта.

Director (распорядитель) - супервизионный класс,
под конролем котрого строитель выполняет скоординированные этапы для создания объекта-продукта.
Распорядитель обычно получает на вход строителя с этапами на выполнение в четком порядке для построения объекта-продукта.
*/

// Product - продукт, который мы строим
type Product struct {
	part1 string
	part2 int
}

// Builder - интерфейс для строителей
type Builder interface {
	BuildPart1()
	BuildPart2()
	GetProduct() Product
}

// ConcreteBuilder1 - конкретная реализация строителя
type ConcreteBuilder1 struct {
	product Product
}

func (b *ConcreteBuilder1) BuildPart1() {
	b.product.part1 = "Part 1 of the product"
}

func (b *ConcreteBuilder1) BuildPart2() {
	b.product.part2 = 42
}

func (b *ConcreteBuilder1) GetProduct() Product {
	return b.product
}

// Director - директор, который управляет строителями
type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct() {
	d.builder.BuildPart1()
	d.builder.BuildPart2()
}

func main() {
	builder := &ConcreteBuilder1{}
	director := NewDirector(builder)

	director.Construct()

	product := builder.GetProduct()

	fmt.Printf("Product built: %+v\n", product)
}
