package pattern

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

	Плюсы паттерна Фабричный метод (Factory Method):

	1. Избавляет клиентский код от необходимости знать конкретные классы создаваемых объектов,
	что уменьшает связанность между классами.
	2. Позволяет подклассам решать, какие объекты создавать, не изменяя общий интерфейс.
	3. Позволяет легко добавлять новые типы продуктов, расширяя фабрику без изменения существующего кода.
	4.Упрощает тестирование, так как можно подменять создаваемые объекты без изменения клиентского кода.

	Минусы паттерна Фабричный метод (Factory Method):

	1. Может привести к созданию большого количества классов фабрик и продуктов, что усложняет структуру программы.
	2. Увеличивает количество классов и интерфейсов, что может затруднить понимание программы.
	3. Паттерн Фабричный метод подходит для ситуаций, когда необходимо делегировать создание объектов
	подклассам, чтобы они могли определить, какие конкретные объекты создавать.
*/

// Product interface
type Product interface {
	Use() string
}

// ConcreteProductA
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using ConcreteProductA"
}

// ConcreteProductB
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using ConcreteProductB"
}

// ProductFactory
type ProductFactory struct{}

func (pf *ProductFactory) CreateProduct(productType string) Product {
	switch productType {
	case "A":
		return &ConcreteProductA{}
	case "B":
		return &ConcreteProductB{}
	default:
		return nil
	}
}

func RunFactoryMethod() {
	factory := &ProductFactory{}

	productA := factory.CreateProduct("A")
	fmt.Println(productA.Use())

	productB := factory.CreateProduct("B")
	fmt.Println(productB.Use())
}
