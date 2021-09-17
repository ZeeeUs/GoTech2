package ownmod

import (
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

// PrintExactTime печатает точное время
func PrintExactTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println("Exact time:", time)
}
