package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"unicode"
)

// Получить текст из файла
func getFileText(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// Разделение текста по словам
func textToWords(s string) []string {
	var words []string // Финальный слайс слов
	var cur []rune     // Слайс типа rune для хранения текущего слова в символах univode

	// Локальная функция для отчистки буфера и добавления слова в финальный слайс
	flush := func() {
		if len(cur) > 0 {
			words = append(words, string(cur))
			cur = cur[:0] // Отчистка буфера
		}
	}

	// Цикл по каждому символу строки
	for _, r := range s {
		// Проверка символа на принадлежность к буквам, если нет, либо игнорируем, либо заканчиваем слово
		switch {
		case unicode.IsLetter(r):
			cur = append(cur, unicode.ToLower(r))
		case unicode.IsSpace(r) || r == ',' || r == '.' || r == ';' || r == '!' || r == '?' || r == ':':
			flush() // Завершение слова
		default:
			// игнорируем символ
		}
	}
	flush() // если последний символ буква

	return words
}

// Создание биграмм
func createBigramms(words []string) map[string][]string {
	// Создание мапы с ключём str и значением в виде слайса
	bigramms := make(map[string][]string)

	// Цикл до предпоследнего слова
	for i := 0; i+1 < len(words); i++ {
		w1, w2 := words[i], words[i+1]          // Текущее и след слово
		bigramms[w1] = append(bigramms[w1], w2) // Добавление след слова к ключу первого
	}
	return bigramms
}

func autocomplete(bigramms map[string][]string, start string, n int) []string {
	start = strings.ToLower(start)
	out := make([]string, 0, n+1) // Инициализация финальной строки
	out = append(out, start)      // добавление первого слова
	curr := start
	for range n {
		current_choise, ok := bigramms[curr] // слайс со словами от текущего + проверка на последнее слово
		if !ok {
			break
		}
		next_word := current_choise[rand.Intn(len(current_choise))]
		out = append(out, next_word)
		curr = next_word
	}
	return out
}

func task11() {
	// Чтение данных из файла
	fileName := "text.txt"
	text := getFileText(fileName)
	fmt.Printf("Исходный текст: %v\n", text)

	// Создание списка слов
	var words []string = textToWords(text)
	fmt.Printf("Слайс слов: %v\n", words)

	// Создание биграмм
	bigramms := createBigramms(words)
	fmt.Printf("Получившиеся биграммы:\n%v\n", bigramms)

	var start_word string
	fmt.Print("Введите слово для генерации продолжения: ")
	fmt.Scan(&start_word)

	// Проверка введенного слова
	if _, ok := bigramms[start_word]; !ok {
		fmt.Println("Заданного слова нет в словаре биграмм")
		return
	}

	// Генерация предложения
	num_auto_words := 4

	seq := autocomplete(bigramms, start_word, num_auto_words)

	fmt.Printf("Сгенерированное предложение:\n%v\n", seq)
}
