package stackv1

import (
	"log"
	"sync"
)

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return &Stack{
		data: make([]int, 0),
	}
}

func (s *Stack) Push(value int) {
	s.data = append(s.data, value)
}

func (s *Stack) Pop() int {
	if len(s.data) == 0 {
		log.Panicln("len of Stack is 0")
	}

	head := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return head
}

func (s *Stack) Head() int {
	return s.data[len(s.data)-1]
}

func populateStack(stck *Stack) {
	for i := range 100 {
		stck.Push(i)
	}
}

func Sample() {
	var mux sync.Mutex
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

			mux.Lock()
			stck.Head()
			stck.Pop()
			mux.Unlock()

		}()
	}

	wg.Wait()
}
