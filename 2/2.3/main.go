package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}
func main() {
	err := Foo()
	fmt.Println(err) // nil
	fmt.Printf("\n%T - Тип хранимого объекта под капотом\n", err)
	fmt.Printf("%v - Значение этого объекта под капотом\n", err)
	fmt.Println("err == nil:", err == nil) // false

	err = nil
	fmt.Printf("\n%T - Тип хранимого объекта под капотом\n", err)
	fmt.Printf("%v - Значение этого объекта под капотом\n", err)
	fmt.Println("err == nil:", err == nil) // true
}