package main

import (
	"fmt"
	"go_course/students"
)

func task12() {
	//Создание группы
	group := students.CreateGroup("Group1")
	fmt.Printf("Создана группа: %q, студентов: %d\n", group.Name, len(group.Students))

	// Загрузка группы из файла
	group2, ok := students.LoadGroup("group.json")
	if !ok {
		fmt.Println("Не удалось загрузить group.json")
		return
	}

	fmt.Printf("Загружена группа: %q, студентов: %d\n", group2.Name, len(group2.Students))
	for _, s := range group2.Students {
		fmt.Printf("  #%d %s %s (GPA=%.1f)\n", s.ID, s.LastName, s.FirstName, s.GPA)
	}

	// Добавление студента к группе
	ok = students.AddStudToGroup(group2, 3, "Петров", "Пётр", 6.0)
	if !ok {
		fmt.Println("Не удалось добавить студента в группу")
		return
	} else {
		fmt.Println("Новый студент Петров Пётр добавлен в группу")
	}

	// Сохранение группы в файл
	if ok := students.SaveGroup(group2, "group_out.json"); !ok {
		fmt.Println("Не удалось сохранить group_out.json")
		return
	}
	fmt.Println("Измененная группа сохранена в group_out.json")

	// Вывод студентов группы
	students_slice := students.GetAllStudList(group2)
	fmt.Println("Список студентов после добавления нового:")
	fmt.Println(students_slice)

	//Вывод студентов по убыванию GPA
	students_rated := students.GetRating(group2)
	fmt.Println("Список студентов по убыванию GPA:")
	fmt.Println(students_rated)

	//Вывод информации о студенте по ID
	fmt.Println("Информация о студенте с ID=2:")
	fmt.Println(students.GetByID(group2, 2))

	// Удаление студента из группы
	students.Del(group2, 2)
	fmt.Println("Спикок студентов после удаления ID=2:")
	fmt.Println(students.GetAllStudList(group2))
}
