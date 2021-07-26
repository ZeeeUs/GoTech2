package time

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func PrintExactTime(){
	time, err := ntp.Time("ntp1.stratum2.ru")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current time: ", time)
}