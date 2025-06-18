package chanv3

import (
	"log"
	"sync"
	"time"
)

/*
	Broadcast single notify.
		+В случае закрытия канала, все читатели разблокируются и прочитают zero value + opened = false.

	NB!: Однако канал нельзя закрыть несколько раз - будет паника, т.к. пытаемся закрыть уже закрытый канал.
	В этом случае используем sync.Cond
*/

func consumer(listen <-chan struct{}) {
	<-listen

	log.Println("i'm notified about event")
}

func broadcast(notify chan<- struct{}) {
	defer close(notify)

	time.Sleep(time.Second * 2)

}

func Sample() {
	var wg sync.WaitGroup
	signals := make(chan struct{})

	wg.Add(3)

	go func() {
		defer wg.Done()

		consumer(signals)
	}()

	go func() {
		defer wg.Done()

		consumer(signals)
	}()

	go func() {
		defer wg.Done()

		broadcast(signals)
	}()

	wg.Wait()

}
