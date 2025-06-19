package lockfree

import (
	"log"
)

/*
	Реализация структуры Stack.
	head -> nil
	push: create item with value: item.next -> head : head -> item.
*/

type item[Value any] struct {
	value any
	next  *item[Value]
}

type Stack[Value any] struct {
	head *item[Value]
}

func NewStack[Value any]() *Stack[Value] {
	return &Stack[Value]{
		head: nil,
	}
}

func (s *Stack[Value]) Push(value Value) {
	// Визуализация проблемы в конкурентной среде:

	// Исходное состояние: s.head = nil

	// Горутина A: Push(1)           Горутина B: Push(2)
	// ├─ Читает s.head = nil        ├─ Читает s.head = nil
	// ├─ Создает item{next:nil}     ├─ Создает item{next:nil}
	// └─ Записывает s.head = item1  └─ Записывает s.head = item2

	// Результат: s.head = item2, item1 потерян!

	s.head = &item[Value]{
		next:  s.head,
		value: value,
	}
}

func (s *Stack[Value]) Pop() Value {
	if s.head == nil {
		log.Panicln("cannot pop nil item")
	}

	value := s.head.value.(Value)
	s.head = s.head.next

	return value
}
