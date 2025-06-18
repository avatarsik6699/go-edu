package rwmuxv1

import (
	"log"
	"sync"
	"time"
)

/*
	RWMutex (shared mutex) - Используется тогда, когда часто читают, но редко пишут.
		дает гарантии:
		- В крит. секции 1 писатель и не более.
		- Не могут присутствовать одновременно читатели и писатели.
		- Читателей может быть много, писатель всегда 1 без читателей.

		g1R -Lock()----Unlock()
		g2W -|---------Lock()-------Unlock()
		g3R -Lock()----Unlock()
											  g4R-|---Lock()
												g5R-|---Lock()
*/

type Counters struct {
	mu sync.RWMutex
	m  map[string]int
}

func (c *Counters) Load(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	v, ok := c.m[key]

	return v, ok
}

func (c *Counters) Store(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	time.Sleep(time.Second * 2)
	c.m[key] = value
}

func Sample() {
	var wg sync.WaitGroup
	wg.Add(6)

	c := Counters{
		m: make(map[string]int),
	}

	go func() {
		defer wg.Done()
		log.Println("start set 444")
		c.Store("some", 444)
	}()

	for i := range 5 {
		go func() {
			defer wg.Done()
			log.Println("try read from store", i)
			v, _ := c.Load("some")
			log.Println("get from store", i, "->", v)

		}()
	}

	wg.Wait()
}
