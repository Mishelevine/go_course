package main

import "fmt"

type Time struct {
	year    int
	month   int
	day     int
	hour    int
	minutes int
}

// Високостный ли год
func isLeap(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return year%4 == 0
}

// Сколько дней в месяце
func daysInMonth(year, month int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if isLeap(year) {
			return 29
		}
		return 28
	default:
		return 30
	}
}

func calculateArrival(departure Time, travel Time) Time {
	var arrival Time

	arrival.year = departure.year
	arrival.month = departure.month
	arrival.day = departure.day
	arrival.hour = departure.hour
	arrival.minutes = departure.minutes

	arrival.minutes += travel.minutes
	carryHour := arrival.minutes / 60
	arrival.minutes %= 60

	arrival.hour += travel.hour + carryHour
	carryDay := arrival.hour / 24
	arrival.hour %= 24

	arrival.day += travel.day + carryDay

	for {
		dim := daysInMonth(arrival.year, arrival.month)
		if arrival.day <= dim {
			break
		}
		arrival.day -= dim
		arrival.month++
		if arrival.month > 12 {
			arrival.month = 1
			arrival.year++
		}
	}

	return arrival
}

func task3() {
	var departure Time
	var travel Time
	fmt.Print("Введите время отправления (день, месяц, год, час, минуты): ")
	fmt.Scan(&departure.day, &departure.month, &departure.year, &departure.hour, &departure.minutes)
	fmt.Print("Введите время пути (часы, минуты): ")
	fmt.Scan(&travel.hour, &travel.minutes)

	arr := calculateArrival(departure, travel)
	fmt.Printf("Ожидаемое время прибытия: %d.%d.%d %d:%d\n",
		arr.day, arr.month, arr.year, arr.hour, arr.minutes)
}
