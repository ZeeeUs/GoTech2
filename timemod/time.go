package timemod

import (
	"fmt"
	"log"

	"github.com/beevik/ntp"
)

func PrintExactTime(){
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current time: ", time)
}
