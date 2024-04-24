package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	s1, s2, s3, s4, s5, s6, s7 := "a4bc2d5e", "abcd", "45", "", "qwe\\4\\5", "qwe\\45", "qwe\\\\5"

	result1, result2, result3, result4, result5, result6, result7 := unpack(s1), unpack(s2), unpack(s3), unpack(s4), unpack(s5), unpack(s6), unpack(s7)

	fmt.Println(result1, result2, result3, result4, result5, result6, result7)
}

func unpack(s string) string {
	var result strings.Builder
	escape := false
	prevChar := ' '
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		if escape {
			result.WriteRune(char)
			prevChar = char
			escape = false
		} else if char == '\\' {
			escape = true
		} else if unicode.IsDigit(char) {
			if prevChar == ' ' {
				return ""
			}
			repeat, err := strconv.Atoi(string(char))
			if err != nil {
				return ""
			}
			result.WriteString(strings.Repeat(string(prevChar), repeat-1))
		} else {
			result.WriteRune(char)
			prevChar = char
		}
	}
	return result.String()
}

// func unpack(s string) (string, error) {
// 	digits := map[byte]bool{
// 		'0': true,
// 		'1': true,
// 		'2': true,
// 		'3': true,
// 		'4': true,
// 		'5': true,
// 		'6': true,
// 		'7': true,
// 		'8': true,
// 		'9': true,
// 	}
// 	var result strings.Builder
// 	var prev byte
// 	escape := false
// 	fmt.Println(s)
// 	for i := 0; i < len(s); i++ {
// 		fmt.Println(int(s[i]), i)
// 		// Проверка на цифру
// 		if digits[s[i]] {
// 			if escape {
// 				result.WriteByte(s[i])
// 			}
// 			if i+1 >= len(s) {
// 				return "", fmt.Errorf("incorrect string lens")
// 			}
// 			if digits[s[i+1]] {
// 				return "", fmt.Errorf("incorrect string i+1")
// 			}

// 			if prev == ' ' {
// 				return "", fmt.Errorf("incorrect string ' '")
// 			}
// 			n, err := strconv.Atoi(string(s[i]))
// 			if err != nil {
// 				return "", err
// 			}
// 			result.WriteString(strings.Repeat(string(prev), n-1))
// 			prev = s[i]
// 			continue
// 		}
// 		// Проверка на \
// 		if escape {
// 			result.WriteString(strings.Repeat(string(s[i]), 2))
// 			escape = false
// 			continue
// 		}
// 		if s[i] == '\\' {
// 			escape = true
// 			continue
// 		}

// 		// Просто доабвление символа
// 		result.WriteByte(s[i])
// 		prev = s[i]
// 		continue
// 	}
// 	return result.String(), nil
// }
