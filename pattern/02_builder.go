package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


	Паттерн строитель позволяет создавать сложные объекты, выполняя пошаговую инициализацию.
	Можно наследуясь от базового билдера создавать другие, с более конкретными конфигурациями.

	Плюсы:
		-Отделение конструирования объекта от его представления: Паттерн Builder позволяет
			создавать сложные объекты пошагово, изолируя процесс конструирования от основного объекта.
		-Упрощение конструирования: Позволяет создавать различные варианты сложных объектов,
			 не усложняя конструктор объекта.
		-Гибкость: Позволяет изменять внутреннюю реализацию объекта без изменения его внешнего интерфейса.
		-Повышение читаемости кода: Последовательное вызывание методов пошагового создания объекта делает
			 код более понятным и читаемым.
	Минусы:
		-Увеличение объема кода: Использование Builder может привести к увеличению объема кода
			 из-за введения дополнительных классов для пошагового конструирования объекта.
		-Сложность поддержки: Для каждого нового типа объекта, требуется создание нового строителя,
			 что может усложнить код и поддержку при большом числе объектов.
		-Недостаточная гибкость: При изменении структуры объекта или добавлении новых свойств,
			 необходимо изменять и строитель, что может привести к сложностям.
		-Усложнение кода для простых объектов: Для создания простых объектов использование
			 Builder может быть излишним и усложнить код.
*/

type FriedEggsBuilderInterface interface {
	EggsCount(value int) FriedEggsBuilderInterface
	Bacon(value bool) FriedEggsBuilderInterface
	Onion(value bool) FriedEggsBuilderInterface

	Build() FriedEggs
}

type FriedEggs struct {
	EggsCount int
	Bacon     bool
	Onion     bool
}

type friedEggsBuilder struct {
	eggsCount int
	bacon     bool
	onion     bool
}

func (f *friedEggsBuilder) Build() FriedEggs {
	return FriedEggs{
		EggsCount: f.eggsCount,
		Bacon:     f.bacon,
		Onion:     f.onion,
	}
}

func newFriedEggsBuilder() FriedEggsBuilderInterface {
	return &friedEggsBuilder{}
}

func (f *friedEggsBuilder) EggsCount(value int) FriedEggsBuilderInterface {
	f.eggsCount = value
	return f
}
func (f *friedEggsBuilder) Bacon(value bool) FriedEggsBuilderInterface {
	f.bacon = value
	return f
}
func (f *friedEggsBuilder) Onion(value bool) FriedEggsBuilderInterface {
	f.onion = value
	return f
}

type cleanFriedEggsBuilder struct {
	friedEggsBuilder
}

func (c *cleanFriedEggsBuilder) Build() FriedEggs {
	return FriedEggs{
		EggsCount: 3,
		Bacon:     false,
		Onion:     false,
	}
}

func newCleanFriedEggsBuilder() FriedEggsBuilderInterface {
	return &cleanFriedEggsBuilder{}
}

func runBuilder() {
	friedEggsBuilder := newFriedEggsBuilder()
	friedEggs := friedEggsBuilder.EggsCount(2).Bacon(true).Onion(true).Build()
	fmt.Println(friedEggs)

	cleanFriedEggsBuilder := newCleanFriedEggsBuilder()
	cleanFriedEggs := cleanFriedEggsBuilder.Build()
	fmt.Println(cleanFriedEggs)
}
