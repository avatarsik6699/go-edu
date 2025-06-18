package condv1

import (
	"log"
	"sync"
)

/*
	sync.Cond - позволяет реализовать засыпание горутин до тех пор, пока явно не придет сигнал о том,
	что можно продолжить выполнение. Например rate limiter, когда нужно ограничить кол-во одновременных
	соединений - максимум 5 одномоментно.

*/

// sync.Cond особенно полезен в следующих случаях:
// 	Реализация пулов ресурсов
// 	Очереди сообщений
// 	Синхронизация параллельных вычислений
// 	Реализация барьеров синхронизации
// 	Управление доступом к ограниченным ресурсам

func Sample() {
	var mux sync.Mutex
	var wg sync.WaitGroup
	var doWork int

	sem := NewSemaphore(5)

	wg.Add(10)

	for range 10 {
		go func() {
			defer wg.Done()

			sem.Acquire()
			log.Println("sem count ->", sem.Available())

			// time.Sleep(time.Second)

			mux.Lock()
			doWork++
			log.Println("do work", doWork)
			mux.Unlock()

			sem.Release()
		}()
	}

	wg.Wait()
}
