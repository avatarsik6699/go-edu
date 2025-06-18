package stackv2

import (
	"log"
	"sync"
)

type Stack struct {
	data []int
	mux  sync.Mutex
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(value int) {
	s.mux.Lock()
	s.data = append(s.data, value)
	s.mux.Unlock()
}

func (s *Stack) Pop() int {
	s.mux.Lock()
	if len(s.data) == 0 {
		log.Panicln("len of Stack is 0")
	}

	head := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	s.mux.Unlock()

	return head
}

func (s *Stack) Head() int {
	s.mux.Lock()
	v := s.data[len(s.data)-1]
	s.mux.Unlock()

	return v
}

func populateStack(stck *Stack) {
	for i := range 100 {
		stck.Push(i)
	}
}

func Sample() {
	// var mux sync.Mutex
	var wg sync.WaitGroup

	stck := NewStack()

	populateStack(stck)

	wg.Add(100)

	for range 100 {
		go func() {
			defer wg.Done()

			/*
				Зачем здесь нужен mux?
				initial: stack [1,2]

				Предполагаю, что если в g1 я прочитал 2 из Head(), то Pop() удалит 2 и запишет [1],
				но это не так.
				g1 ---Head()-> 2 |-------------------------Pop() -> []
				g2 --------------Head() -> 2,Pop() -> [1] |

				Solve:
					1. Обернуть обе операции в доп. mux извне.
					2. Соединить 2 операции в одну синхронизированную изнутри.
			*/

			// mux.Lock()
			stck.Head()
			// Здесь все равно может встроиться горутина?
			// Хотя каждая операция (Head() и Pop()) по отдельности защищена мьютексом,
			// атомарность между ними не гарантирована. Это классический пример "check-then-act" race condition.
			// Для исправления нужно либо:
			// 	Создать новую синхронизированную операцию, объединяющую Head() и Pop()
			//	Использовать внешний мьютекс для защиты последовательности операций
			stck.Pop()
			// mux.Unlock()

		}()
	}

	wg.Wait()
}
