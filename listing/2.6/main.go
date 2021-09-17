package main

import (
	"fmt"
)

func main() {
	var s = []string{"pattern", "listing", "develop"}
	modifySlice(s)
	fmt.Println(s) // develop listing develop
}

func modifySlice(i []string) {
	i[0] = "develop"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}

// Слайс в Go - ссылочный тип
// В функцию мы передаём по значению. Когда меняем в первый раз, то мы по прежнему работаем с тем же слайсом, который в main
// когда происходит операция append у нас уже появляется локальная копия этого слайса, но уже с другим размером и дальше мы работаем с ним
