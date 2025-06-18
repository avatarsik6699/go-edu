package chanv2

import (
	"log"
	"sync"
	"time"
)

/*
	Channel single notify
*/

func consumer(listen <-chan struct{}) {
	<-listen

	log.Println("i'm notified about event")
}

func producer(notify chan<- struct{}) {
	time.Sleep(time.Second * 2)

	notify <- struct{}{}
}

func Sample() {
	var wg sync.WaitGroup
	signals := make(chan struct{})

	wg.Add(2)

	go func() {
		defer wg.Done()

		consumer(signals)
	}()

	go func() {
		defer wg.Done()

		producer(signals)
	}()

	wg.Wait()

}
