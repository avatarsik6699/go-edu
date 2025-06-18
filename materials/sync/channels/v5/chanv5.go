package chanv5

import (
	"log"
	"time"
)

/*
	CHANNELS INTERNAL.

	make(chan any) - возвращает указатель на структуру hchan.

	closed поле равно 32bit - why?
		У atomic.Bool под капотом используется uint32 значение.
*/

func Sample() {
	ch := make(chan int)

	go func() {
		log.Println("start")
		<-ch
		log.Println("end")
	}()

	time.Sleep(time.Second)
	log.Println("close")
	close(ch)

	time.Sleep(100 * time.Millisecond)
}
