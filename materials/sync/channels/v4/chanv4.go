package chanv4

import (
	"log"
	"time"
)

/*
	Buff/Unbuff channels.

	Кто должен закрывать канал?
		- тот, кто пишет в канал.
		- если писателей несколько, то тот, кто создавал писателей.

	Что если не закрыть канал?
		- произойдет утечка горутин HANGING GORUTINES (подвисшие горутины).
*/

func canGetLenAndCapOfChannels() {
	syncExplChan := make(chan int)
	syncImplChan := make(chan int, 0)
	buffChan := make(chan int, 5)

	// close(buffChan)
	// close(syncExplChan)
	// close(syncImplChan)

	buffChan <- 1
	buffChan <- 2
	buffChan <- 3
	buffChan <- 4
	buffChan <- 5
	// buffChan <- 6 deadlock

	log.Println(
		len(syncExplChan), cap(syncExplChan),
		len(syncImplChan), cap(syncImplChan),
		len(buffChan), cap(buffChan),
	)
}

func readFromBuffChAfterClose() {
	buffChan := make(chan int, 5)

	buffChan <- 1
	buffChan <- 2
	buffChan <- 3
	buffChan <- 4
	buffChan <- 5

	firstValue := <-buffChan  // 1
	secondValue := <-buffChan // 2
	log.Println(firstValue, secondValue)
	close(buffChan)

	// После закрытия выгребаются оставшиеся данные

	firstValue = <-buffChan  // 3
	secondValue = <-buffChan // 4
	log.Println(firstValue, secondValue)
	lastValue := <-buffChan // 5
	log.Println(lastValue)
	boundedValue := <-buffChan // zero value
	log.Println(boundedValue)
	boundedValue = <-buffChan // zero value
	log.Println(boundedValue)
}

func compareChannels() {
	ch1 := make(chan int, 0)
	ch2 := make(chan int, 0)

	ch3 := ch1
	log.Println(ch3 == ch1)
	log.Println(ch3 == ch2)
}

func goleakv1() {
	ch := make(chan int)

	go func() {
		for v := range ch {
			log.Println(v)
		}

		// Код до сюда не дойдет и горутина не освободится, пока канал не будет закрыт. - утечка.
		log.Println("exit from gorutine")
	}()

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch)

	time.Sleep(time.Second)
	log.Println("Done")
}

func goleakv2() <-chan int {
	// first-response-win strategy
	ch := make(chan int)

	for i := range 10 {
		go func() {
			// 1. Можно решить проблему 9 подвешеных каналов через буфф. канал,
			// Когда остальные горутины запишут свои данные в оставшиеся ячейки и освободятся
			// Сборщик мусора потом удалит канал с данными после завершения программы.

			// 2. Другой вариант через select с таймаутами и default.

			// 3. Закрыть канал в доп. горутине + wg.
			ch <- i // 9 каналов останутся висеть
		}()
	}

	return ch
}

func prioritizationReadFromChannels() int {

	lowWeightChannel := make(chan int, 1)
	highWeightChannel := make(chan int, 1)

	lowWeightChannel <- 1
	go func() {
		time.Sleep(time.Second * 1)
		highWeightChannel <- 2
	}()

	currentValue := 0

loop:
	for {
		select {
		case v, ok := <-lowWeightChannel:
			log.Println("run 1", v, ok)
			currentValue = v
			if ok {
				close(lowWeightChannel)
			}
			// lowWeightChannel = nil
			continue
		case v := <-highWeightChannel:
			log.Println("run 2")
			currentValue = v

			break loop
		case <-time.After(time.Second * 2):
			log.Println("run timer")
			break loop
		}
	}

	return currentValue
}

func readFromNilCh() {
	ch := make(chan int, 1)

	ch = nil
	ch <- 1   // nil канал не может принять данные - deadlock
	v := <-ch // nil канал не может отдать данные

	// Решение: Всегда проверяйте, что канал не nil перед использованием,
	// или используйте select с default для неблокирующих операций.

	log.Println(v)
}

func Sample() {
	readFromNilCh()

}
