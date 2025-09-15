package main

import "fmt"

type Time struct {
	year    int
	month   int
	day     int
	hour    int
	minutes int
}

func calculateArrival(departure Time, travel_time Time) {
	var arrival Time
	minutes := departure.minutes + travel_time.minutes
	if minutes >= 60 {
		arrival.minutes = minutes % 60
		travel_time.hour += minutes / 60
	} else {
		arrival.minutes = minutes
	}

	hours := departure.hour + travel_time.hour
	if hours >= 24 {
		arrival.hour = hours % 24
		travel_time.day += hours / 24
	} else {
		arrival.hour = hours
	}

	days := departure.day + travel_time.day
	if days > 30 {
		arrival.day = days % 30
		if arrival.day == 0 {
			arrival.day = 30
			travel_time.month += days/30 - 1
		} else {
			travel_time.month += days / 30
		}
	} else {
		arrival.day = days
	}

	months := departure.month + travel_time.month
	if months > 12 {
		arrival.month = months % 12
		if arrival.month == 0 {
			arrival.month = 12
			travel_time.year += months/12 - 1
		} else {
			travel_time.year += months / 12
		}
	} else {
		arrival.month = months
	}

	years := departure.year + travel_time.year
	arrival.year = years

	fmt.Printf("Ожидаемое время прибытия: %d.%d.%d %d:%d\n", arrival.day, arrival.month, arrival.year, arrival.hour, arrival.minutes)
}

func task3() {
	var departure Time
	var travel_time Time
	fmt.Print("Введите время отправления (день, месяц, год, час, минуты): ")
	fmt.Scan(&departure.day, &departure.month, &departure.year, &departure.hour, &departure.minutes)
	fmt.Print("Введите время пути (часы, минуты): ")
	fmt.Scan(&travel_time.hour, &travel_time.minutes)

	calculateArrival(departure, travel_time)
}
