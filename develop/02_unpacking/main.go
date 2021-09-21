package main

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

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Unpacking сканирует и распаковывает строки
func Unpacking(s string) string{
	var sc scanner.Scanner
	var res string
	var prev string
	sc.Init(strings.NewReader(s))
	sc.Mode = scanner.ScanChars | scanner.ScanInts
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		switch tok {
		case scanner.Int:
			num, err := strconv.Atoi(sc.TokenText())
			if err != nil {
				fmt.Printf("Something went wrong %s", err)
			}
			res += strings.Repeat(prev, num-1)
		default:
			prev = sc.TokenText()
			res += sc.TokenText()
		}
	}
	return res
}

func main() {
	fmt.Println(Unpacking("a4bc2d5e"))
}
