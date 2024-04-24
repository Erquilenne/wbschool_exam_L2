package pattern

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern

	Плюсы паттерна цепочка вызовов (Chain of Responsibility):

	1. Уменьшает зависимость между отправителем запроса и получателем, что позволяет гибко настраивать
	и изменять порядок обработки запросов.
	2. Позволяет добавлять новые обработчики или изменять их порядок без изменения клиентского кода.
	3. Позволяет реализовать различные варианты обработки запроса, например, в зависимости
	 от типа запроса или других условий.

	Минусы паттерна цепочка вызовов (Chain of Responsibility):

	1. Не гарантирует, что запрос будет обработан, если он достигнет конца цепочки без выполнения.
	2. Может усложнить отладку из-за неявной обработки запроса и разветвленной логики обработчиков.

	Паттерн цепочка вызовов подходит для ситуаций, когда имеется несколько объектов,
	способных обработать запрос, и клиент не знает заранее, какой объект сможет обработать запрос.
*/

import "fmt"

// Handler interface
type Handler interface {
	SetNext(Handler)
	Handle(int) bool
}

// ConcreteHandler
type ConcreteHandler struct {
	Next Handler
}

func (h *ConcreteHandler) SetNext(next Handler) {
	h.Next = next
}

func (h *ConcreteHandler) Handle(number int) bool {
	if number%3 == 0 && number%5 == 0 {
		fmt.Println("FizzBuzz")
		return true
	} else if number%3 == 0 {
		fmt.Println("Fizz")
		return true
	} else if number%5 == 0 {
		fmt.Println("Buzz")
		return true
	}

	if h.Next != nil {
		return h.Next.Handle(number)
	}

	return false
}

func RunChain() {
	handler1 := &ConcreteHandler{}
	handler2 := &ConcreteHandler{}
	handler3 := &ConcreteHandler{}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	for i := 1; i <= 30; i++ {
		handler1.Handle(i)
	}
}
