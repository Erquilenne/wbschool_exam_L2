package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fieldsPtr := flag.String("f", "", "choose fields (columns)")
	delimiterPtr := flag.String("d", "\t", "use another delimiter")
	separatedPtr := flag.Bool("s", false, "only print lines with the delimiter")
	flag.Parse()

	var lines []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	selectedColumns := getSelectedColumns(lines, *fieldsPtr, *delimiterPtr, *separatedPtr)
	for _, selectedColumn := range selectedColumns {
		fmt.Println("selectedColumn", selectedColumn)
	}
}

func getSelectedColumns(lines []string, fields string, delimiter string, separated bool) []string {
	var selectedColumns []string
	for _, line := range lines {
		if separated && !strings.Contains(line, string(delimiter)) {
			continue
		}
		columns := strings.Split(line, delimiter)
		if fields != "" {
			fileds := strings.Split(fields, ",")
			selectedLine := []string{}
			for _, field := range fileds {
				fieldNum, _ := strconv.Atoi(field)
				if fieldNum > len(columns) {
					continue
				}
				selectedLine = append(selectedLine, columns[fieldNum-1])
				continue
			}
			joinedLine := strings.Join(selectedLine, delimiter)
			if joinedLine == "" {
				continue
			}
			selectedColumns = append(selectedColumns, joinedLine)
		} else {
			selectedColumns = append(selectedColumns, line)
		}

	}
	return selectedColumns
}
