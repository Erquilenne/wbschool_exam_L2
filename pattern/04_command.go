package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern

	Плюсы паттерна Команда (Command):

	1. Инкапсулирует запрос как объект, что позволяет параметризовать клиентские объекты с
	отличающимися запросами.
	2. Упрощает добавление новых команд и изменение существующей логики без изменения самих команд.
	3. Позволяет создавать очереди команд, отменять операции, а также регистрировать запросы.
	4. Повышает расширяемость кода и уменьшает связанность между отправителем запроса и получателем.

	Минусы паттерна Команда (Command):

	1. Может привести к созданию большого количества классов команд, что усложняет структуру программы.
	2. Может увеличить объем кода из-за создания отдельных классов для каждой команды.
	3. Возможно усложнение понимания логики программы из-за разделения запроса и его обработки на различные объекты.
	4. Паттерн Команда подходит для ситуаций, когда необходимо инкапсулировать запрос как объект
	и передавать его как аргумент, а также для реализации отмены операций, регистрации запросов
	и управления очередью команд.
*/

// Command interface
type Command interface {
	Execute()
}

// ConcreteCommand
type ConcreteCommand struct {
	Receiver *Receiver
}

func (c *ConcreteCommand) Execute() {
	c.Receiver.Action()
}

// Receiver
type Receiver struct{}

func (r *Receiver) Action() {
	fmt.Println("Receiver executing action")
}

// Invoker
type Invoker struct {
	Command Command
}

func (i *Invoker) ExecuteCommand() {
	i.Command.Execute()
}

func RunVisitor() {
	receiver := &Receiver{}
	command := &ConcreteCommand{Receiver: receiver}
	invoker := &Invoker{Command: command}

	invoker.ExecuteCommand()
}
