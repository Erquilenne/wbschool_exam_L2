package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern

	Плюсы паттерна Состояние (State):

	1. Упрощение кода: Паттерн Состояние позволяет вынести поведение, связанное с определенным состоянием,
	в отдельные классы, что упрощает управление кодом.
	2. Четкое разграничение состояний: Позволяет явно определить различные состояния объекта и их
	взаимодействие, что улучшает читаемость кода.
	3. Гибкость и расширяемость: Обеспечивает гибкость в добавлении новых состояний и изменении
	поведения объекта в зависимости от состояния.
	4. Упрощенное тестирование: Позволяет тестировать каждое состояние объекта независимо друг от
	друга, что улучшает тестирование и обнаружение ошибок.

	Минусы паттерна Состояние (State):

	1. Увеличение числа классов: Использование паттерна Состояние может привести к увеличению числа классов,
	 особенно при большом числе состояний, что может усложнить структуру программы.
	2. Повышение сложности конфигурации: Требуется правильная настройка и управление состояниями,
	 что может быть сложным в случае наличия множества состояний и их переходов.


	 Паттерн Состояние обычно применяется в ситуациях, когда объект может иметь различное поведение
	 в зависимости от своего текущего состояния, и когда необходимо обеспечить гибкость и легкость
	 сопровождения кода.
*/

import "fmt"

// State interface
type State interface {
	DoAction(context *StateContext)
}

// ConcreteStateA
type ConcreteStateA struct{}

func (s *ConcreteStateA) DoAction(context *StateContext) {
	fmt.Println("State A: Performing action")
	context.state = &ConcreteStateB{}
}

// ConcreteStateB
type ConcreteStateB struct{}

func (s *ConcreteStateB) DoAction(context *StateContext) {
	fmt.Println("State B: Performing action")
	context.state = &ConcreteStateA{}
}

// Context
type StateContext struct {
	state State
}

func (c *StateContext) Request() {
	c.state.DoAction(c)
}

func RunState() {
	context := &StateContext{state: &ConcreteStateA{}}

	context.Request()
	context.Request()
}
