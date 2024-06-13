package main

import (
	"fmt"
)

type stack struct {
	s    []any // слайс в котором хранятся значения в стеке
	head int   // индекс головы стека
}

func newStack(size int) *stack {
	return &stack{
		s:    make([]any, size),
		head: -1,
	}
}

// push - добавление в стек значения
func push(s *stack, v any) {
	if s.head < len(s.s)-1 {
		s.head++
		s.s[s.head] = v
	} else {
		fmt.Println("Stack overflow")
	}
}

// pop - получение значения из стека и его удаление с вершины
func pop(s *stack) any {
	if s.head >= 0 {
		v := s.s[s.head]
		s.s[s.head] = nil // Обнуляем значение на вершине стека (опционально)
		s.head--
		return v
	}
	fmt.Println("Stack underflow")
	return nil
}

// peek - просмотр значения на вершине стека
func peek(s *stack) any {
	if s.head >= 0 {
		return s.s[s.head]
	}
	fmt.Println("Stack is empty")
	return nil
}

func main() {
	st := newStack(5)

	push(st, 10)
	push(st, 20)
	push(st, 30)

	fmt.Println("Peek:", peek(st)) // Ожидается 30

	fmt.Println("Pop:", pop(st)) // Ожидается 30
	fmt.Println("Pop:", pop(st)) // Ожидается 20

	push(st, 40)
	fmt.Println("Peek:", peek(st)) // Ожидается 40

	fmt.Println("Pop:", pop(st)) // Ожидается 40
	fmt.Println("Pop:", pop(st)) // Ожидается 10
	fmt.Println("Pop:", pop(st)) // Ожидается "Stack underflow" и nil
}
