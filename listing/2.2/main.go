package main

import "fmt"

func test() (x int){
	defer func() {
		x++
	}()
	x = 1
	return
}
func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}
func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}

// Программа выведет 2 1
// Ключевое слово defer добавляет функцию в стек, т.е. функции исполняются
// по методу FIFO (First In First Out)

