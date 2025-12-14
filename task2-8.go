package main

import (
	"fmt"
)

type LimitedSizeStack struct {
	max_len int   // максимальная вместимость стека
	buf     []int // кольцевой буфер
	head    int   // индекс самого глубокого элемента в buf
	count   int   // сколько элементов сейчас лежит в стеке
}

// Конструктор стека
func NewLimitedSizeStack(maxLen int) LimitedSizeStack {
	return LimitedSizeStack{buf: make([]int, maxLen), max_len: maxLen}
}

func (s *LimitedSizeStack) Push(item int) {
	// стек не полный
	if s.count < s.max_len {
		i := (s.head + s.count) % s.max_len // индекс для следующего элемента (делим чтобы не вылезал за пределы)
		s.buf[i] = item
		s.count++ // увеличиваем count
		return
	}

	// стек полный
	s.buf[s.head] = item              //заменяем самый старый элемент на новый
	s.head = (s.head + 1) % s.max_len // передвигаем индекс самого глубокого
}

func (s *LimitedSizeStack) Pop() (int, bool) {
	// в стеке не осталось элементов
	if s.count == 0 {
		return 0, false
	}

	top_i := (s.head + s.count - 1) % s.max_len // индекс верхнего элемента
	v := s.buf[top_i]                           // получаем значение
	s.count--
	return v, true
}

// получение текущего размера
func (s *LimitedSizeStack) Count() int {
	return s.count
}

func task2_8() {
	stack := NewLimitedSizeStack(2)

	fmt.Println("Кладем 10")
	stack.Push(10) // в стеке 10
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Кладем 20")
	stack.Push(20) // в стеке 10, 20
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Кладем 30")
	stack.Push(30) // в стеке 20, 30 (удален самый глубокий элемент 10)
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Кладем 40")
	stack.Push(40) // в стеке 30, 40 (удален 20)
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Вызываем Pop")
	fmt.Println(stack.Pop()) // достает 40, в стеке остаётся 30
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Кладем 50")
	stack.Push(50) // в стеке 30, 50
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Вызываем Pop")
	fmt.Println(stack.Pop()) // достает 50, в стеке остаётся 30
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Вызываем Pop")
	fmt.Println(stack.Pop()) // достает 30, стек остаётся пуст
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())

	fmt.Println("Вызываем Pop")
	fmt.Println(stack.Pop())
	fmt.Printf("Сейчас в стеке %v элементов\n", stack.Count())
}
