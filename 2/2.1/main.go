package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [5]int{76, 77, 78, 79, 80} // массив
	var b []int = a[1:4] // b инициализируеься как слайс, в данном случае тип можно не указывать
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(b))

}
