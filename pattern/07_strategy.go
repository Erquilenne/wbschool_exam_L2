package pattern

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern

	Плюсы паттерна Стратегия (Strategy):

	1. Разделение ответственности: Паттерн Стратегия позволяет разделить ответственность за выполнение
	определенной функциональности на разные классы, каждый из которых реализует свою собственную стратегию.
	 Это улучшает модульность и переиспользуемость кода.
	2. Обеспечение открытости для расширения: Стратегии могут быть легко добавлены или изменены без
	изменения классов, использующих стратегии. Это упрощает расширение функциональности.
	3. Изолированное тестирование: Стратегии могут быть протестированы независимо от других компонентов
	 системы, что упрощает разработку и отладку.
	4. Простая замена стратегии: Можно легко заменить одну стратегию на другую без изменения
	клиентского кода, что упрощает настройку и изменение логики системы.

	Минусы паттерна Стратегия (Strategy):

	1. Усложнение структуры программы: Паттерн Стратегия может привести к увеличению числа классов в
	 программе из-за создания отдельных стратегий, что может усложнить структуру программы.
 	2. Увеличение сложности конфигурации: При использовании паттерна Стратегия необходимо правильно
	настроить контекст и выбрать подходящую стратегию для выполнения задачи, что может быть сложно в
	случае большого числа стратегий.
	3. Повышенное использование памяти: Каждый дополнительный класс стратегии потребляет определенное
	 количество памяти, что может привести к увеличению использования памяти программой.
*/

import "fmt"

// Strategy interface
type Strategy interface {
	ExecuteStrategy(int, int) int
}

// AddStrategy
type AddStrategy struct{}

func (s *AddStrategy) ExecuteStrategy(a, b int) int {
	return a + b
}

// SubtractStrategy
type SubtractStrategy struct{}

func (s *SubtractStrategy) ExecuteStrategy(a, b int) int {
	return a - b
}

// Context
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.ExecuteStrategy(a, b)
}

func RunStrategy() {
	context := &Context{}

	addStrategy := &AddStrategy{}
	context.SetStrategy(addStrategy)
	result1 := context.ExecuteStrategy(5, 3)
	fmt.Println("Result of addition:", result1)

	subtractStrategy := &SubtractStrategy{}
	context.SetStrategy(subtractStrategy)
	result2 := context.ExecuteStrategy(5, 3)
	fmt.Println("Result of subtraction:", result2)
}
