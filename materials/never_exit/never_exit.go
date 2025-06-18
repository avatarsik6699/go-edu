package neverexit

import (
	"fmt"
	"log"
	"time"
)

func TaskWithPanic() {
	for {
		time.Sleep(time.Second)
		panic("unexpected situation")
		// fmt.Println("some work...")
	}
}

func Task() {
	for {
		time.Sleep(time.Second)
		fmt.Println("some work...")
	}
}

func NeverExit(name string, fn func()) {
	defer func() {
		log.Println("defer call")
		// if v := recover(); v != nil {
		// 	log.Println(name, "is crashed - restarting...", v)
		// }

	}()

	if fn != nil {
		fn()
	}
}
