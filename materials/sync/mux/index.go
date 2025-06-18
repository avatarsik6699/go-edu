package main

import (
	"log"
	"sync"
	"time"
)

/*
	Mutex - синхронизирует конкурентный доступ к участку памяти.
	Критическая секция - участок памяти "под защитой" мьютекса
	(тут может находиться только 1 горутина).

	*Нужно в основном тогда, когда происходят мутабельные операции с разделяемым значением.

	Global/Local gourtines queues - в этих очередях находятся горутины в статусе ready (
	ждут когда их поставят на выполнение ).

	- На уровне ОС mutex блокирует поток. У каждого mux есть свой wait_list (linked list), где хранятся
	заблокированные потоки. То есть каждый mux хранит список заблоченых им потоков.

	- На уровне go runtime блокируется gorutine с выносом в WaitQueue. Поток не блокируется и
	может взять другую горутину на исполнение. Архитектурно имеется глобальное дерево, где
	адрес каждого мьютекса сопоставляется с заблокированной горутиной.
		mux в runtime Go может работать в 2ух режимах: Normal и Starvation mode.

			Normal mode -> горутины хранятся в fifo очереди и следующей захватывает мьютекс первая
			в очереди горутина.

			Starvation mode -> механизм предотвращения голодания,
			где долго ожидающие горутины получают приоритет (сама горутина, которая разлочит mux передаст
			"эстафету" первой горутине в очереди).

	NB:
		1. При этом нет гарантии какая горутина успеет захватить mux, это происходит
		конкурентно.
		2. 1 mux может блокировать n секций кода. Причем если одна из секций захвачена, то
		все остальные секции также считаются захваченными.
*/

const (
	tasksCount = 10000
)

var (
	mux sync.Mutex
)

func main() {
	var sum int

	wg := ParallelTasks(tasksCount, func(idx int) {
		log.Println("run", idx)
		time.Sleep(time.Second)
		mux.Lock()
		sum += 1
		mux.Unlock()
	})

	log.Println("wait all tasks")
	wg.Wait()
	log.Println("all tasks has completed", sum)
}

/*
Хочу параллельно запускать tasksNums и дожидаться их исполнения
с получением результата.
*/
func ParallelTasks(tasksNums int, task func(int)) *sync.WaitGroup {
	log.Println("start ParallelTasks")

	var wg sync.WaitGroup
	wg.Add(tasksNums)

	for i := range tasksNums {
		go func(idx int) {
			defer wg.Done()

			task(idx)
		}(i)
	}

	log.Println("return ParallelTasks")
	return &wg
}
