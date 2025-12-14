package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	ErrNegative = errors.New("число не может быть отрицательным")
	ErrTooLarge = errors.New("число больше максимально допустимого для множества")
)

// интерфейс
type Set interface {
	Add(x int) error
	Del(x int)
	Check(x int) bool
	Cap() int
}

// множество как неэкспортируемый массив
type ListSet struct {
	Set
	maxN  int
	items []int
}

// создание множества
func NewListSet(maxN int) (Set, error) {
	if maxN < 0 {
		return nil, fmt.Errorf("maxN должен быть >= 0")
	}
	return &ListSet{maxN: maxN, items: make([]int, 0)}, nil
}

// добавление элемента
func (s *ListSet) Add(x int) error {
	if x < 0 { // отрицательное
		return ErrNegative
	}
	if x > s.maxN { // больше допустимого
		return ErrTooLarge
	}

	// проверка на существование элемента
	for _, v := range s.items {
		if v == x {
			return nil
		}
	}

	s.items = append(s.items, x)
	return nil
}

// удаление элемента
func (s *ListSet) Del(x int) {
	// проверка на отрицательное или больше допустимого
	if x < 0 || x > s.maxN {
		return
	}

	// проходим по массиву и удаляем элемент
	for i, v := range s.items {
		if v == x {
			// меняем текущий на последний и обрезаем до -1
			s.items[i] = s.items[len(s.items)-1]
			s.items = s.items[:len(s.items)-1]
			return
		}
	}
}

// проверка наличия элемента
func (s *ListSet) Check(x int) bool {
	// проверка на отрицательное или больше допустимого
	if x < 0 || x > s.maxN {
		return false
	}

	// ищем элемент
	for _, v := range s.items {
		if v == x {
			return true
		}
	}
	return false
}

// Возвращает максимальное число для множества
func (s *ListSet) Cap() int {
	return s.maxN
}

func (s *ListSet) String() string {
	if len(s.items) == 0 {
		return "{}"
	}

	// создание копии и сортировка
	tmp := make([]int, len(s.items))
	copy(tmp, s.items)
	sort.Ints(tmp)

	var b strings.Builder

	b.WriteString("{")
	for i, v := range tmp {
		if i > 0 {
			b.WriteString(", ")
		}
		b.WriteString(strconv.Itoa(v))
	}
	b.WriteString("}")
	return b.String()
}

// множество как неэкспортируемый логический массив
type SimpleSet struct {
	Set
	maxN int
	m    []bool
}

func NewSimpleSet(maxN int) (Set, error) {
	if maxN < 0 {
		return nil, fmt.Errorf("maxN должен быть >= 0")
	}
	return &SimpleSet{maxN: maxN, m: make([]bool, maxN+1)}, nil
}

// добавление элемента
func (s *SimpleSet) Add(x int) error {
	if x < 0 { // отрицательное
		return ErrNegative
	}
	if x > s.maxN { // больше допустимого
		return ErrTooLarge
	}

	// индекс массива равен элементу, поэтому просто вставляем true
	s.m[x] = true

	return nil
}

// удаление элемента
func (s *SimpleSet) Del(x int) {
	// проверка на отрицательное или больше допустимого
	if x < 0 || x > s.maxN {
		return
	}

	// подставляем false
	s.m[x] = false
}

// проверка на наличие элемента
func (s *SimpleSet) Check(x int) bool {
	// проверка на отрицательное или больше допустимого
	if x < 0 || x > s.maxN {
		return false
	}

	return s.m[x]
}

// Возвращает максимальное число для множества
func (s *SimpleSet) Cap() int { return s.maxN }

func (s *SimpleSet) String() string {
	first := true

	var b strings.Builder

	b.WriteString("{")
	for i := 0; i <= s.maxN; i++ {
		if s.m[i] {
			if !first {
				b.WriteString(", ")
			}
			first = false
			b.WriteString(strconv.Itoa(i))
		}
	}
	b.WriteString("}")
	return b.String()
}

// Универсальная функция Fill
func Fill(s Set, input string) {
	// записываем элементы в слайс
	num := strings.Fields(input)

	// проход по элементам и обработка ощибок
	for _, n := range num {
		x, _ := strconv.Atoi(n)

		addErr := s.Add(x)

		if addErr != nil {
			fmt.Println("Ошибка:", addErr)
		}
	}
}

func task2_9() {
	var set_var int
	var maxN int
	fmt.Println("Выберите реализацию множества:")
	fmt.Println("1 - ListSet (массив элементов)")
	fmt.Println("2 - SimpleSet (логический массив)")
	fmt.Print("Введите 1 или 2: ")
	fmt.Scan(&set_var)

	fmt.Print("Введите maxN для множества: ")
	fmt.Scan(&maxN)

	var s Set
	var err error

	switch set_var {
	case 1:
		s, err = NewListSet(maxN)
	case 2:
		s, err = NewSimpleSet(maxN)
	default:
		fmt.Println("Некорректное значение варианта множества")
		return
	}

	if err != nil {
		fmt.Println("Ошибка создания множества:", err)
		return
	}

	// удаляем возможные \n
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	fmt.Println("Введите элементы множества через пробел (одной строкой):")

	// читаем строку
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	Fill(s, line)

	var choose int

	for {
		fmt.Println("\nМеню:")
		fmt.Println("1 - Добавить элемент")
		fmt.Println("2 - Удалить элемент")
		fmt.Println("3 - Проверить наличие элемента")
		fmt.Println("4 - Вывести множество")
		fmt.Println("5 - Вывести максимально возможный элемент множества ")
		fmt.Println("0 - Выход")

		fmt.Print("Выберите пункт: ")
		fmt.Scan(&choose)

		var x int

		switch choose {
		case 1:
			fmt.Print("Введите число для добавления: ")
			fmt.Scan(&x)
			err := s.Add(x)

			if err != nil {
				fmt.Println("Ошибка добавления элемента:", err)
			}
		case 2:
			fmt.Print("Введите число для удаления: ")
			fmt.Scan(&x)
			s.Del(x)
		case 3:
			fmt.Print("Введите число для проверки на наличие: ")
			fmt.Scan(&x)
			if s.Check(x) {
				fmt.Printf("Число %v присутствует в множестве\n", x)
			} else {
				fmt.Printf("Число %v не присутствует в множестве\n", x)
			}
		case 4:
			fmt.Printf("Текущее множество: \n %v", s)
		case 5:
			fmt.Printf("Максимально возможный элемент множества: \n %v", s.Cap())
		case 0:
			return
		default:
			fmt.Println("Неизвестная команда")
		}
	}
}
