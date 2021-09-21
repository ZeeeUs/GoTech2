package main

import (
	"fmt"
	"github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/app/apiserver"
)

func main() {
	config := apiserver.NewConfig()
	fmt.Println(config)
}
