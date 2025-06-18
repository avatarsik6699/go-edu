package oncev1

import (
	"log"
	"sync"
)

/*
	Once - используется тогда, когда во множестве горутин нужно выполнить действие только 1 раз.

	Patterns:
		1. Lazy initialization
		2. Sync singleton
*/

type Single struct{}

var instance *Single
var once sync.Once

func GetInstance() *Single {
	// Выполнится только 1 раз
	once.Do(func() {
		instance = &Single{}
	})

	return instance
}

func Sample() {
	var wg sync.WaitGroup
	var once sync.Once

	action := func() {
		log.Println("action run")
	}

	wg.Add(10)

	for range 10 {
		go func() {
			defer wg.Done()
			once.Do(action)
		}()
	}

	wg.Wait()
}
