package lockfree

import (
	"log"
	"sync"
	"time"
)

/*
LOCK FREE - неблокирующая синхронизация. Дает доступ к объектам без блокировок.

NB!: почитать про ABA Problem. - когда при конкурентном доступе по адресу очищается значение, но
параллельно другая структура читает этот адрес. В это время первая записывает другое значение в адрес.
В итоге вторая читает не то значение, которое ожидала.

Паттерны lock free структур/алгоритмов.
1. RCU( read-copy-update ) - подход, где синхронизация построена на атомиках.
*/
type SecretStruct struct {
	privateField string
}

func Sample() {
	stack := NewStack[int]()
	var wg sync.WaitGroup

	wg.Add(10)

	for range 5 {
		go func() {
			time.Sleep(time.Millisecond)
			defer wg.Done()
			v := stack.Pop()
			log.Println(v)
		}()
	}

	for i := range 5 {
		go func() {
			defer wg.Done()
			stack.Push(i)
		}()
	}

	wg.Wait()

	log.Println("done")
}
