package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	anagramDictionary := []string{"пятак", "листок", "кулон", "столик", "пятка", "тяпка", "слиток", "клоун", "уклон"}
	anagramms := findAnagramGroups(&anagramDictionary)
	fmt.Println(anagramms)
}

func findAnagramGroups(words *[]string) map[string][]string {
	anagramGroups := make(map[string][]string)

	for _, word := range *words {
		word = strings.ToLower(word)
		// Выравниваем буквы в алфавитном порядке
		sortedWord := sortString(word)

		anagramGroups[sortedWord] = append(anagramGroups[sortedWord], word)
	}

	// Убираем группы с 1 элементом
	for key, value := range anagramGroups {
		if len(value) == 1 {
			delete(anagramGroups, key)
		}
	}

	// Меняем ключи с алфавитных на первое слово из среза
	namedGroups := make(map[string][]string, len(anagramGroups))
	for key, value := range anagramGroups {
		namedGroups[value[0]] = value
		delete(anagramGroups, key)
	}

	return namedGroups
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
