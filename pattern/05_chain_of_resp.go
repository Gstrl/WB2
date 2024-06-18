package main

import "fmt"

/*
Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Применимость паттерна "цепочка вызовов":
1. Обработка различных типов запросов различными способами, когда не известны типы запросов и их последовательность
2. Выполнение нескольких обработчиков в определённом порядке

Плюсы паттерна "цепочка вызовов":
1. Принцип единой ответственности: Можете отделить классы, вызывающие операции, от классов, выполняющих операции.
2. Принцип открытости/закрытости: Можно вводить в приложение новые обработчики, не ломая существующий код.
3. Контроль порядка обработки запросов

Минусы паттерна "цепочка вызовов":
1. Некоторые запросы могут быть не обработаны
*/

// IHandler создаем интерфейс определяющий сигнатуру функций элементов цепочки
type IHandler interface {
	SetNext(IHandler) IHandler
	Handle(string) string
}

type Handler struct {
	next IHandler
}

func (h *Handler) SetNext(next IHandler) IHandler {
	h.next = next
	return next
}

func (h *Handler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

type FirstReceiver struct {
	Handler
}

func (f *FirstReceiver) Handle(request string) string {
	// this one will not handle any requests
	return f.next.Handle(request)
}

func (f *FirstReceiver) setNext(next IHandler) IHandler {
	f.next = next
	return next
}

type SecondReceiver struct {
	Handler
}

func (s *SecondReceiver) Handle(request string) string {
	if request == "Hello" {
		return "World"
	}
	return s.next.Handle(request)
}

func (s *SecondReceiver) setNext(next IHandler) IHandler {
	s.next = next
	return next
}

func main() {
	FirstReceiver := &FirstReceiver{}
	SecondReceiver := &SecondReceiver{}

	FirstReceiver.setNext(SecondReceiver)

	request := "Hello"
	response := FirstReceiver.Handle(request)
	fmt.Println(response)

}
