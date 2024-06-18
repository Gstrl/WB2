package main

import "fmt"

/*
Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,
а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Применимость паттерна "фасад":
1. Когда система становится сложной и требует упрощения интерфейса для клиентов.
2. Когда необходимо скрыть детали реализации подсистем от клиентов.

Плюсы паттерна "фасад":
1. Упрощение интерфейса для клиентов.
2. Сокрытие сложности системы.
3. Соблюдение принципа единственной ответственности: Фасад отвечает за координацию работы подсистем, оставляя им отдельные области ответственности

Минусы паттерна "фасад":
1. Может ввести в систему дополнительные слои абстракции, что усложнит обслуживание в будущем.
2. Может привести к созданию god object.
*/

//В данном примере мы создаем фасад Facade,
//который скрывает сложность взаимодействия с подсистемой Subsystem1
//и Subsystem2. При вызове метода Operation() фасад инициализирует подсистемы
//и выполняет действия с их помощью, предоставляя удобный интерфейс для клиента.

// Subsystem1 сложная подсистема
type Subsystem1 struct{}

func (s *Subsystem1) Method1() {
	fmt.Println("Subsystem1: Method1")
}

func (s *Subsystem1) Method2() {
	fmt.Println("Subsystem1: Method2")
}

// Subsystem2 еще одна сложная подсистема
type Subsystem2 struct{}

func (s *Subsystem2) Method1() {
	fmt.Println("Subsystem2: Method1")
}

func (s *Subsystem2) Method2() {
	fmt.Println("Subsystem2: Method2")
}

// Facade Фасад, скрывающий сложность подсистемы
type Facade struct {
	subsystem1 *Subsystem1
	subsystem2 *Subsystem2
}

func NewFacade() Facade {
	return Facade{
		subsystem1: &Subsystem1{},
		subsystem2: &Subsystem2{},
	}
}

func (f *Facade) Operation() {
	fmt.Println("Facade initializes subsystems:")
	f.subsystem1.Method1()
	fmt.Println("Facade orders subsystems to perform the action:")
	f.subsystem1.Method2()
}

func main() {
	facade := NewFacade()
	facade.Operation()
}
