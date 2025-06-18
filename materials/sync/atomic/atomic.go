package atomicv1

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
)

/*

	TRUE / FALSE SHARING -
	True/False sharing (или cache line sharing) - это проблема производительности, возникающая при работе с атомиками в многопоточной среде.
	True sharing происходит, когда несколько процессоров активно модифицируют одну и ту же кэш-линию (обычно 64 байта). Это приводит к постоянной синхронизации кэшей между процессорами, что значительно снижает производительность.
	False sharing возникает, когда разные процессоры работают с разными переменными, но эти переменные находятся в одной кэш-линии. Даже если процессоры не взаимодействуют напрямую, они вынуждены синхронизировать кэши из-за близкого расположения данных в памяти.
	Для предотвращения false sharing в Go часто используют padding (дополнение) структур до размера кэш-линии:

	type PaddedAtomic struct {
    value atomic.Int64
    _     [56]byte // padding для предотвращения false sharing
}
*/

var data map[string]string
var isInit atomic.Bool

func initializeV1() {
	if !isInit.Load() {
		/*
			Здесь могут втиснуться множество горутин и привести к тому,
			что инициализация выполнится несколько раз.

			Операция чтения isInit.Load и записи isInit.Store хоть сами по себе и атомарны, но
			они не синхронизированы. Необходима атомарная операция чтения+запись - реализовано в v2.
		*/
		runtime.Gosched()

		isInit.Store(true)
		data = make(map[string]string)
		log.Println("init")
	}
}

func initializeV2() {
	/*
		CAS - атомарная операция чтения+запись.
		Проверяет равно ли текущее значение old (false), если да, то меняет на new (true)
		и возвращает true, т.к. операция выполнена успешно.
	*/
	if isInit.CompareAndSwap(false, true) {

		runtime.Gosched()

		data = make(map[string]string)
		log.Println("init")
	}
}

func Sample() {
	var wg sync.WaitGroup
	wg.Add(100)

	for range 100 {
		go func() {
			defer wg.Done()
			initializeV2()
		}()
	}

	wg.Wait()
}
