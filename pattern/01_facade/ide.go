package main

import "fmt"

type TextEditor struct {}

func (te TextEditor) writeCode(){
	fmt.Println("Написание кода")
}

func (te TextEditor) saveCode() {
	fmt.Println("Сохранение кода")
}

type Compiler struct {}

func (cp Compiler) Compile(){
	fmt.Println("Компиляция кода")
}

type CLR struct {}

func (clr CLR) Execute(){
	fmt.Println("Выполнение приложения")
}

func (clr CLR) Finish() {
	fmt.Println("Завершение приложения")
}
