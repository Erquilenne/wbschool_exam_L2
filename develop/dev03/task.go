package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Необходимо указать название файла")
		return
	}

	fileName := args[len(args)-1]
	var lines []string

	file, err := os.OpenFile(fileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	key := flag.Int("k", 0, "указание колонки для сортировки")
	numeric := flag.Bool("n", false, "сортировать по числовому значению")
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	lines = sortLines(lines, *key, *numeric, *reverse, *unique)

	fmt.Println("after sorting - ", lines)

	for _, line := range lines {
		fmt.Println(line)
	}
}

func sortLines(lines []string, key int, numeric bool, reverse bool, unique bool) []string {

	if numeric {
		sort.SliceStable(lines, func(i, j int) bool {
			a, b := lines[i], lines[j]
			if key > 0 {
				aSplit, bSplit := splitLine(a), splitLine(b)
				a, b = aSplit[key-1], bSplit[key-1]
			}
			numA, errA := strconv.Atoi(a)
			numB, errB := strconv.Atoi(b)
			fmt.Println("numa =", numA, "numb =", numB, "a =", a, "b =", b)
			if errA == nil && errB == nil {
				if reverse {
					return numA > numB
				}
			}
			return numA < numB
		})
	} else {
		if key > 0 {
			sort.SliceStable(lines, func(i, j int) bool {
				a, b := lines[i], lines[j]
				aSplit, bSplit := splitLine(a), splitLine(b)
				a, b = aSplit[key-1], bSplit[key-1]
				return a < b
			})
		}
		if reverse {
			sort.SliceStable(lines, func(i, j int) bool {
				a, b := lines[i], lines[j]
				return a < b
			})
		}
	}

	if unique {
		lines = removeDuplicates(lines)
	}
	return lines
}
func splitLine(line string) []string {
	// Простой сплитер по пробелам
	return strings.Fields(line)
}

func removeDuplicates(lines []string) []string {
	encountered := map[string]bool{}
	result := []string{}

	for _, line := range lines {
		if !encountered[line] {
			encountered[line] = true
			result = append(result, line)
		}
	}

	return result
}
