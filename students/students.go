package students

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Student struct {
	ID        int     `json:"id"`
	LastName  string  `json:"last_name"`
	FirstName string  `json:"first_name"`
	GPA       float32 `json:"gpa"`
}

type Group struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

func CreateGroup(name string) *Group {
	return &Group{
		Name:     name,
		Students: make([]Student, 0, 16),
	}
}

// Загрузка группы их json файла
func LoadGroup(fileName string) (*Group, bool) {
	// Чтение файла
	data, err := os.ReadFile(fileName)
	if err != nil {
		return &Group{}, false
	}

	group := Group{}

	// парсинг полученного json
	err2 := json.Unmarshal(data, &group)
	if err2 != nil {
		return &Group{}, false
	}

	if group.Students == nil {
		group.Students = make([]Student, 0)
	}

	if group.Name == "" {
		return &Group{}, false
	}

	return &group, true
}

// Сохранение группу в json файл
func SaveGroup(group *Group, fileName string) bool {
	if group == nil {
		return false
	}

	// Создание json c форматированием
	data, err := json.MarshalIndent(group, "", "  ")
	if err != nil {
		return false
	}
	// Попытка записи в файл
	err2 := os.WriteFile(fileName, data, 0644)
	return err2 == nil
}

// Добавление студента в группу
func AddStudToGroup(group *Group, id int, surname string, name string, gpa float32) bool {
	if group == nil {
		return false
	}

	// Проверка уникальности id
	for i := range group.Students {
		if group.Students[i].ID == id {
			return false
		}
	}

	// Добавление студента
	group.Students = append(group.Students, Student{
		ID:        id,
		LastName:  surname,
		FirstName: name,
		GPA:       gpa,
	})
	return true
}

// Формирование отсортированного списка студентов группы
func GetAllStudList(group *Group) []string {
	// Проверка на существование студентов
	if group == nil || len(group.Students) == 0 {
		return []string{}
	}

	out := make([]string, 0, len(group.Students))

	// Проходимся по всем студентам и добавляем строку с ИФ в ответ
	for _, s := range group.Students {
		out = append(out, s.LastName+" "+s.FirstName)
	}

	// Сортировка слайса
	sort.Strings(out)

	return out
}

// Cписок студентов в порядке убывания GPA
func GetRating(group *Group) []string {
	// Проверка на существование студентов
	if group == nil || len(group.Students) == 0 {
		return []string{}
	}

	// Копируем слайс, чтобы сортировать его
	sorted := make([]Student, len(group.Students))
	copy(sorted, group.Students)

	// Сортировка по убыванию GPA
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].GPA > sorted[j].GPA
	})

	// Формирование финального слайса
	out := make([]string, 0, len(sorted))
	for _, s := range sorted {
		out = append(out, fmt.Sprintf("%d %s %s %.1f", s.ID, s.LastName, s.FirstName, s.GPA))
	}
	return out
}

// Получение студента по id
func GetByID(group *Group, id int) string {
	// Проверка на существование группы
	if group == nil {
		return ""
	}

	// Проходим по всем студентам и ищем нужный id
	for _, s := range group.Students {
		if s.ID == id {
			return fmt.Sprintf("%d %s %s %.2f", s.ID, s.LastName, s.FirstName, s.GPA)
		}
	}

	// Если студент не найден - пустая строка
	return ""
}

// Удаление студента из группы
func Del(group *Group, id int) {
	// Проверка на существование группы и студентов
	if group == nil || len(group.Students) == 0 {
		return
	}

	// Проходим по мтудентам и ищем ID
	for i, s := range group.Students {
		if s.ID == id {
			// Копируем элементы с i+1 и вставляем их с текущего элемента
			copy(group.Students[i:], group.Students[i+1:])
			// Удалаем последний элемент (копия предпоследнего)
			group.Students = group.Students[:len(group.Students)-1]
			return
		}
	}
}
