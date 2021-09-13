package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

// Функция unpacking сканирует и распаковывает строки
func unpacking(s string) string{
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

func escapeseq(s string) {
	var res string
	var arr []string
	var buf []string
	//s += `\`
	arr = strings.Split(s, "")
	for i, v := range arr {
		if v == `\` || arr[i+1] == "\n"{
			if arr[i+1] == `\`{
				res += v
			}
			buf = append(buf, res)
			res =""
			continue
		}
		res += v
	}
	res = ""
	for _, v := range buf{
		res += unpacking(v)
	}
	fmt.Println(res)
}

func main() {
	fmt.Println(unpacking("a4bc2d5e"))
	//unpacking("abcd")
	//unpacking("")
	//unpacking("3245")
	//unpacking("3eb5d5")
	//tdt(`qwe\a4\r5`)
	//tdt(`qwe\\5`)
	//escapeseq(`qwe\\5`)
	//escapeseq(`qwe\a4\r5`)
	//escapeseq(`qwe\4\5`)
}
