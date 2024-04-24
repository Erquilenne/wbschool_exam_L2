package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	options := make(map[string]interface{})

	after := flag.Int("A", 0, "Print +N lines after matching line")
	before := flag.Int("B", 0, "Print +N lines before matching line")
	context := flag.Int("C", 0, "Print ±N lines around matching line")
	count := flag.Bool("c", false, "Print only a count of matching lines")
	ignoreCase := flag.Bool("i", false, "Ignore case when matching")
	invert := flag.Bool("v", false, "Invert the match, exclude matching lines")
	fixed := flag.Bool("F", false, "Match fixed string, not pattern")
	lineNum := flag.Bool("n", false, "Print line numbers")

	flag.Parse()

	options["after"] = *after
	options["before"] = *before
	options["context"] = *context
	options["fixed"] = *fixed
	options["ignoreCase"] = *ignoreCase
	options["invert"] = *invert

	args := flag.Args()
	if len(args) < 2 {
		log.Fatal("Usage: go run main.go [flags] pattern filename")
	}

	pattern := args[0]
	filename := args[1]

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(content), "\n")

	matchingLines, addedLines := findMatches(lines, pattern, options)

	if *count {
		fmt.Println(len(matchingLines))
	} else {
		for _, line := range addedLines {
			if *lineNum {
				if strings.Contains(matchingLines[line], pattern) {
					fmt.Println(fmt.Sprintf("%d:%s", line+1, matchingLines[line]))
				} else {
					fmt.Println(fmt.Sprintf("%d-%s", line+1, matchingLines[line]))
				}
			} else {
				fmt.Println(matchingLines[line])
			}
		}
	}
}

func findMatches(lines []string, pattern string, options map[string]interface{}) (map[int]string, []int) {
	ignoreCase := options["ignoreCase"].(bool)
	fixed := options["fixed"].(bool)
	invert := options["invert"].(bool)
	before := options["before"].(int)
	after := options["after"].(int)
	context := options["context"].(int)

	matchingLines := make(map[int]string, 0)

	addedLines := []int{}
	for i, line := range lines {
		match := false
		if ignoreCase {
			match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
		} else if fixed {
			match = strings.EqualFold(line, pattern)
		} else {
			match = strings.Contains(line, pattern)
		}

		if (invert && !match) || (!invert && match) {
			if before > 0 || context > 0 {
				var add int
				if before > context {
					add = before
				} else {
					add = context
				}
				addBefore(lines, add, i, &matchingLines, &addedLines)
			}
			matchingLines[i] = line
			addedLines = append(addedLines, i)
			if after > 0 || context > 0 {
				var add int
				if after > context {
					add = after
				} else {
					add = context
				}
				addAfter(lines, add, i, &matchingLines, &addedLines)
			}
		}
	}
	addedLines = removeDuplicates(addedLines)
	return matchingLines, addedLines
}

func addBefore(lines []string, add int, i int, matchingLines *map[int]string, addedLines *[]int) {
	if i-add < 0 {
		for y := 0; y < i; y++ {
			added := false
			for _, num := range *addedLines {
				if num == y {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[y] = lines[y]
				*addedLines = append(*addedLines, y)
			}
		}
	} else {
		sequence := getSequence(i, i-add, false)
		for _, j := range sequence {
			added := false
			for _, num := range *addedLines {
				if num == j {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[j] = lines[j]
				*addedLines = append(*addedLines, j)
			}
		}
	}
}

func addAfter(lines []string, add int, i int, matchingLines *map[int]string, addedLines *[]int) {
	if i+add >= len(lines) {
		for y := i; y < len(lines); y++ {
			added := false
			for _, num := range *addedLines {
				if num == y {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[y] = lines[y]
				*addedLines = append(*addedLines, y)
			}
		}
	} else {
		sequence := getSequence(i, i+add, true)
		for _, j := range sequence {
			added := false
			for _, num := range *addedLines {
				if num == j {
					added = true
					break
				}
			}
			if !added {
				(*matchingLines)[j] = lines[j]
				*addedLines = append(*addedLines, j)
			}
		}
	}
}

func getSequence(start int, finish int, up bool) []int {
	sequence := []int{}
	if up {
		for i := start; i <= finish; i++ {
			sequence = append(sequence, i)
		}
	} else {
		for i := start; i >= finish; i-- {
			sequence = append(sequence, i)
		}
	}
	return sequence
}

func removeDuplicates(slice []int) []int {
	unique := make(map[int]bool)
	result := []int{}

	for _, item := range slice {
		if !unique[item] {
			unique[item] = true
			result = append(result, item)
		}
	}

	sort.Ints(result)

	return result
}
