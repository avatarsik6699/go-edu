package memorybarriers

import (
	"log"
	"runtime"
	"sync/atomic"
)

var a string
var done bool

func setup() {

	/*
		Данная секция может быть переупорядочена компилятором для оптимизиации, т.к.
		этот код для данного "потока" выполняется идентично в любом порядке.
			a = "Hello, world"
			done = true

			done = true
			a = "Hello, world"


		Как запретить переупорядочивание инструкций?
			- используются логические барьеры памяти LoadLoad, StoreStore, LoadStore, StoreLoad и тд
			которые в Go реализованы через atomiс, скрывая абстракцией то, какой на самом деле барьер
			там используется (возможно зависит от арх.).

		NB!: мьютексы используют барьеры памяти на уровне асм команд.
	*/
	// --------------
	a = "Hello, world"
	done = true
	// --------------

	if done {
		log.Println(len(a)) // expected 12
	}

}

var av2 string
var donev2 atomic.Bool

func setupWithBarrier() {
	av2 = "Hello, world"
	donev2.Store(true)

	if donev2.Load() {
		log.Println(len(av2))
	}
}

func Sample() {
	go setupWithBarrier()

	for !donev2.Load() {
		// Переключение контекста на другую горутину для уменьшения кол-ва итераций вхолостую.
		runtime.Gosched()
	}

	log.Println(av2) // expected "Hello, world"
}
