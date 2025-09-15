package main

import (
	"fmt"
)

const taskMessage = `Укажите номер задания:
1. Форматирование оценки из 10-ти бальной в 5-ти бальную
2. Сумма цифр в числе

!Для выхода введите 0`

func chooseTask() {
	for {
		var task_num int
		fmt.Println(taskMessage)

		fmt.Scan(&task_num)

		switch task_num {
		case 0:
			return
		case 1:
			task1()
		case 2:
			task2()

		default:
			fmt.Println("Ошибка. Неверный номер задания")
		}
	}
}

func main() {
	chooseTask()
}
