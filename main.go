package main

import (
	"fmt"
)

const taskMessage = `Укажите номер задания:
1. Форматирование оценки из 10-ти бальной в 5-ти бальную
2. Сумма цифр в числе
3. Посчитать время прибытия
4. Функция Эйлера
5. Банковский вклад
6. Редукция матрицы
7. Сортировка строк матрицы
8. Пересечение мультимножеств
9. Генератор перестановок
10. AA-дерево
11. Интеллектуальный помощник ввода

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
		case 3:
			task3()
		case 4:
			task4()
		case 5:
			task5()
		case 6:
			task6()
		case 7:
			task7()
		case 8:
			task8()
		case 9:
			task9()
		case 10:
			task10()
		case 11:
			task11()
		case 12:
			task12()
		default:
			fmt.Println("Ошибка. Неверный номер задания")
		}
	}
}

func main() {
	chooseTask()
}
