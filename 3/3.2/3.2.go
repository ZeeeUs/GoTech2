package main

import (
	"fmt"
	"strconv"
	"strings"
)

// unpacker распаковывает строки
func unpacker(s string) {
	// Разбираем строку и записываем посимвольно в слайс
	arr := strings.Split(s, "")
	var str string
	var counter int

	for i, v := range arr {
		// Пытаемся перевести элемент слайса в int
		num, err := strconv.Atoi(v)
		if err != nil {
			str += v
			counter++
			continue
		}

		if i != 0 {
			// Проверка на наличие букв в слайсе
			if counter == 0 {
				continue
			}
			str += strings.Repeat(arr[i-1], num - 1)
		}
	}

	fmt.Println(str)
}

func main() {
	unpacker("a4bc2d5e")
	unpacker("abcd")
	unpacker("34")
	unpacker("34a3")
	unpacker("")
	unpacker("234f4a")
}
