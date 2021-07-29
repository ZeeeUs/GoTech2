package OwnMod

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func PrintExactTime() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	fmt.Println("Exact time:", time)
}