package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern


	Паттерн фасад нужен для объединения в одной структуре нескольких
	модулей с нужными методами, что бы можно было из одного места вызывать
	их нужные методы

	Плюсы:

		-Упрощение интерфейса: Фасад предоставляет простой интерфейс для работы с сложной
		 	системой, что упрощает использование этой системы клиентским кодом.
		-Сокрытие сложности: Фасад скрывает сложность внутренней структуры и взаимодействия
		 	компонентов системы, предоставляя клиентам только необходимый функционал.
		-Уменьшение зависимостей: Клиентский код взаимодействует только с фасадом,
		 	что позволяет уменьшить зависимости между клиентом и внутренними компонентами системы.
		-Улучшение читаемости кода: Использование фасада делает код более читаемым,
		 	так как клиентский код вызывает только один метод фасада вместо непосредственного взаимодействия с различными компонентами системы.
	Минусы:

		-Ограничение гибкости: Фасад может скрывать некоторые возможности внутренних
			компонентов системы, что может привести к ограничению гибкости в настройке и
			расширении функционала.
		-Добавление еще одного уровня абстракции: Использование фасада может привести
			к добавлению еще одного уровня абстракции, что может усложнить понимание и
			поддержание системы.
		-Возможность нежелательного расширения: Фасад может стать местом, где клиенты
			начинают добавлять нежелательную логику, что противоречит его цели упрощения
			интерфейса.
*/

type span struct{}

func (s *span) Fry(ingredients map[string]string) {
	if ingredients["egg"] == "raw" {
		ingredients["egg"] = "fried"
	}
	if ingredients["onion"] == "chopped" {
		ingredients["onion"] = "fried"
	}
	if ingredients["bacon"] == "raw" {
		ingredients["bacon"] = "fried"
	}
}

func (s *span) CheckDone(ingredients map[string]string) bool {
	if ingredients["egg"] == "fried" && ingredients["onion"] == "fried" && ingredients["bacon"] == "fried" {
		return true
	} else {
		return false
	}
}

type ingredients struct{}

func (i *ingredients) GetFriedEggsIngredients() map[string]string {
	ingredients := map[string]string{
		"egg":   "raw",
		"onion": "raw",
		"bacon": "raw",
	}

	return ingredients
}

type knife struct{}

func (k *knife) Chop(ingredients map[string]string) {
	if ingredients["onion"] == "raw" {
		ingredients["onion"] = "chopped"
	}
}

type friedEggsFacade struct {
	span        span
	ingredients ingredients
	knife       knife
}

func (f *friedEggsFacade) FryEggs() {
	ingredients := f.ingredients.GetFriedEggsIngredients()
	f.knife.Chop(ingredients)
	f.span.Fry(ingredients)

	if f.span.CheckDone(ingredients) {
		fmt.Println("Fried eggs are ready!")
	}
}

func main() {

	friedEggs := friedEggsFacade{
		span:        span{},
		ingredients: ingredients{},
		knife:       knife{},
	}
	friedEggs.FryEggs()
}
