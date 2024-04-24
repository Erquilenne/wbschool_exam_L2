package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern

	 Плюсы паттерна Посетитель:
	 1. Разделение закольцевания: Посетитель позволяет разделить логику, связанную с обходом элементов,
 	от самих элементов. Это улучшает модульность и делает код более гибким.
 	2. Упрощение расширения: Добавление нового посетителя позволяет легко расширить функциональность элементов
 	без изменения самих элементов.
 	3. Разделение ответственности: Посетитель отделяет логику обхода элементов от самих элементов, упрощая их реализацию.
 	4. Удобство для новых операций: Добавление новой операции к элементам может быть выполнено, добавив новый посетитель,
 	без изменения самих элементов.

	Минусы паттерна Посетитель:
 	1. Сложность: Паттерн Посетитель может увеличить сложность кода из-за введения дополнительных интерфейсов и классов.
 	2. Нарушение инкапсуляции: Посетитель может нарушить инкапсуляцию элементов, поскольку он требует, чтобы элементы
 	предоставляли метод Accept, что может быть не всегда желательно.
 	3. Ограниченная поддержка некоторых языков: В некоторых языках программирования, особенно с динамической типизацией,
 	реализация паттерна Посетитель может быть более сложной из-за необходимости работы с типами во время выполнения.
 	4. Усложнение добавления новых элементов: Добавление нового типа элемента требует изменений во всех посетителях,
 	что может привести к усложнениям и необходимости внесения изменений во множество мест кода.
*/

// Интерфейс посетителя
type Visitor interface {
	Visit(element Element)
}

// Конкретный посетитель
type ConcreteVisitor struct{}

func (v *ConcreteVisitor) Visit(element Element) {
	switch element.(type) {
	case *ConcreteElementA:
		element.(*ConcreteElementA).OperationA()
	case *ConcreteElementB:
		element.(*ConcreteElementB).OperationB()
	}
}

// Интерфейс элемента
type Element interface {
	Accept(visitor Visitor)
}

// Конкретный элемент A
type ConcreteElementA struct{}

func (e *ConcreteElementA) Accept(visitor Visitor) {
	visitor.Visit(e)
}

func (e *ConcreteElementA) OperationA() {
	fmt.Println("OperationA() of ConcreteElementA")
}

// Конкретный элемент B
type ConcreteElementB struct{}

func (e *ConcreteElementB) Accept(visitor Visitor) {
	visitor.Visit(e)
}

func (e *ConcreteElementB) OperationB() {
	fmt.Println("OperationB() of ConcreteElementB")
}

func runVisitor() {
	// Создание элементов
	elementA := &ConcreteElementA{}
	elementB := &ConcreteElementB{}

	// Создание посетителя
	visitor := &ConcreteVisitor{}

	// Посещение элементов посетителем
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
