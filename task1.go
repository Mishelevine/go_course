package main

import "fmt"

func translateHseMark(mark int) string {
	if mark < 4 {
		return "Неудовлетворительно"
	} else if mark < 6 {
		return "Удовлетворительно"
	} else if mark < 8 {
		return "Хорошо"
	} else {
		return "Отлично"
	}
}

func task1() {
	var mark int
	fmt.Println("Введите оценку в 10-ти бальной системе: ")
	fmt.Scan(&mark)
	fmt.Printf("Оценка в 5-ти бальной системе: %s\n", translateHseMark(mark))
}
