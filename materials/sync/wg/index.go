package main

import (
	"log"
	"sync"
)

/*
	WaitGroup - ждет окончание нескольких параллельных операций.
	Похоже на Promise.all?.

	g1 -------|done()
	g2 --|done()
	g3 ----|done()
	wg.Wait()-------| Все задачи закончены, двигаемся дальше.

	NB:
		1. Нужно копировать только по ссылке.
		2. done() лучше вызывать в defer, чтобы точно просигнализировать о том, что
		горутина закончила выполнение, даже если упала с паникой.
		3. Вероятно стоит комбинировать с чем-то ещё (таймеры), чтоб не заблокировать ожидание надолго
		если одна из задач исполняется слишком долго.
*/

const (
	tasksCount = 10
)

func main() {

	wg := ParallelTasks(tasksCount, func(idx int) {
		log.Println("run", idx)
	})

	log.Println("wait all tasks")
	wg.Wait()
	log.Println("all tasks has completed")
}

// 2025/06/17 10:59:03 start ParallelTasks
// 2025/06/17 10:59:03 run 3
// 2025/06/17 10:59:03 return ParallelTasks
// 2025/06/17 10:59:03 run 0
// 2025/06/17 10:59:03 run 1
// 2025/06/17 10:59:03 run 2
// 2025/06/17 10:59:03 wait all tasks
// 2025/06/17 10:59:03 run 7
// 2025/06/17 10:59:03 run 9
// 2025/06/17 10:59:03 run 6
// 2025/06/17 10:59:03 run 8
// 2025/06/17 10:59:03 run 4
// 2025/06/17 10:59:03 run 5
// 2025/06/17 10:59:03 all tasks has completed

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
