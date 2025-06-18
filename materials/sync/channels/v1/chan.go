package chanv1

import (
	"log"
	"time"
)

func Sample() {
	data := make(chan int)

	go func() {
		defer close(data)
		for i := range 4 {
			time.Sleep(time.Second)
			data <- i + 1
		}
	}()

	for {
		var value int
		isChanOpen := true

		select {
		default:
			log.Println("fast end", isChanOpen)
			continue
		case value, isChanOpen = <-data:
			log.Println(value, isChanOpen)
			if value == 2 {
				continue
			} else if value == 3 {
				break
			}

			if !isChanOpen {
				return
			}
		}

		log.Println(value)
	}
}
