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
12. Студенты
13. Знакомство с горутинами
14. Сумма массива
15. Перемножение матриц
16. Конвейерная обработка
17. Гипотеза 3n+1
18. Считаем зарплату
19. Логистическая компания
20. LimitedSizeStack
21. Множества

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
		case 13:
			task2_1()
		case 14:
			task2_2()
		case 15:
			task2_3()
		case 16:
			task2_4()
		case 17:
			task2_5()
		case 18:
			task2_6()
		case 19:
			task2_7()
		case 20:
			task2_8()
		case 21:
			task9()

		default:
			fmt.Println("Ошибка. Неверный номер задания")
		}
	}
}

func main() {
	chooseTask()
}
