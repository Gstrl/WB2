package main

import "fmt"

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
