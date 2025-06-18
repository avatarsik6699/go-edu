package main

import (
	"log"
	"time"
)

// Generator pattern - не блокирует инструкции в теле функции.
// func generator() <-chan int {
// 	wg := &sync.WaitGroup{}
// 	ch := make(chan int)

// 	wg.Add(2)

/**
В данном случае горутины выполняются недетерминировано - планировщик решает
кого запускать первым (но чаще всего запускается 2 горутина lifo + fifo?).

В данном случае будет ошибка panic из-за того, что одна горутина уже закрыла канал,
а другая пытается в него писать (пишет в закрытый канал).
*/

// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i <= 5; i++ {
// 			log.Println("go1: ready to write: ", i+1)
// 			time.Sleep(time.Millisecond * 100) // добавить эту строку
// 			ch <- i + 1
// 		}
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		for i := 5; i <= 10; i++ {
// 			log.Println("go2: ready to write: ", i+1)
// 			ch <- i + 1
// 		}

// 	}()

// 	go func() {
// 		log.Println("wait for to end of write/read to/from ch")
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	return ch
// }

/*
	Task: написать 3 функции
		generator - генерит числа от 1 до 10.
		doubler - умножает числа на 2, имитируя работу (500ms).
		reader - читает и выводит числа на экран
*/

func main() {
	// -- Non buffered channels --

	/**
	Non-Initialized Channel - nil.

	Неинизиализированный канал (содержит просто nil)
	Когда вы пытаетесь записать.прочитать nil канал,
	операция блокируется навсегда, потому что:
	- Нет буфера []buff для хранения данных
	- Нет горутин, которые могли бы прочитать данные
	- Нет механизма для обработки такой операции

	var nonInitializedChannel chan any
	close(nonInitializedChannel) panic
	<- nonInitializedChannel (nil) deadlock
	nonInitializedChannel (nil) <- 1 deadlock
	*/

	/**
	Initialized channel - make(chan int)

	Правильное решение - всегда инициализировать канал перед использованием:
	nonBuferedInitializedChannel := make(chan int)

	// Чтение из закрытого канала возвращает nil значение типа (в данном случае 0).
	// Чтение / запись в канал + отсутствие свободных горутин - deadlock.
	<-nonBuferedInitializedChannel deadlock
	nonBuferedInitializedChannel <- 1 deadlock
	*/

	// Pipline pattern - ch1 <- ch2 <- ch3
	// reader(doubler(generator(10)))

	/*
		Оператор Select является блокирующим. Используется в ситуации, когда
		нужно прочитать ИЛИ из канала А, ИЛИ из канала В и т.д.

		Как избежать deadlock?
		1. default: позволяет выйти из deadlock, если заблокированы все case.
		2. timer.After === timer.Timer - по истечение времени пишет в канал и закрывает
		3. close(ch)
		4. ctx Done()
	*/

	ch := make(chan int)

	go func() {
		log.Println("start go")
		time.Sleep(2000);
		ch <- 1;
		log.Println("try to write new value go")
	}()

}

func randomTimeWorker(num time.Duration) {
	time.Sleep(num * time.Second)
}

// Если переданная функция выполняется более 3 сек -> отмена
func withTimeout[F func()](fn F) {
	timer := time.NewTimer(time.Second * 3)
	ch := make(chan any)

	go func() {
		defer func() {
			timer.Stop()
			time.Sleep(time.Second * 2)
			close(ch)
		}()
		fn()
	}()

	select {
	case <-ch:
	case <-timer.C:
		panic("time exceeded")
	}
}

/** ---------------------------------------------- */

func reader(chr <-chan int) {

	for v := range chr {
		log.Println(v)
	}

}

func doubler(chr <-chan int) <-chan int {
	chw := make(chan int)

	go func() {
		defer close(chw)
		for number := range chr {
			time.Sleep(time.Second / 2)

			chw <- number * 2
		}
	}()

	return chw
}

func generator(n int) <-chan int {
	chw := make(chan int)

	go func() {
		defer close(chw)
		for i := range n {
			chw <- i + 1
		}
	}()

	return chw
}
