package main

import (
	"github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/config"
	"github.com/ZeeeUs/GoTech2/develop/11_HTTPServer/simpleApp"
	"log"
)

func main() {
	config := config.NewDefConfig()

	s := simpleapp.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
