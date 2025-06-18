package mapv1

import (
	"log"
	"sync"
)

/*
	CACHE CONTENTION - ситуация, когда потоки читают/пишут в ячейку памяти в одной и той же кеш-линии,
	что ведет к конфликтам в кеше (то есть ядра мешают друг другу).

	- Может возникнуть при использовании RWMutex, т.к. он хранит в себе разделяемый потоками
	счетчик кол-ва читаемых горутин, из-за чего кеш линия постоянно инвалидируется для всех потоков
	и ядер. Чем > горутин, тем > cache contention

	Solved: Synchronized Map.
*/

type Counters struct {
	m sync.Map
}

func (c *Counters) Get(key string) int {
	value, _ := c.m.Load(key)

	if intValue, ok := value.(int); ok {
		return intValue
	} else {
		panic("internal error")
	}
}
func (c *Counters) Set(key string, value int) {
	c.m.Store(key, value)
}

func Sample() {
	var wg sync.WaitGroup
	wg.Add(6)

	c := Counters{}

	for i := range 5 {
		go func() {
			defer wg.Done()
			log.Println("try read from store", i)
			v := c.Get("some")
			log.Println("get from store", i, "->", v)

		}()
	}

	go func() {
		defer wg.Done()
		log.Println("start set 444")
		c.Set("some", 444)
	}()

	wg.Wait()
}
