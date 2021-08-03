package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Функция unpacking сканирует и распаковывает строки
func unpacking(s string) {
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
	fmt.Println(res)
}

func main() {
	//unpacking("a4bc2d5e")
	//unpacking("abcd")
	//unpacking("")
	//unpacking("3245")
	unpacking("3eb5d5")
}
