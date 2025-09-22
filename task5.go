package main

import (
	"fmt"
)

func task5() {
	startSum := 0.0
	percent := 0.0
	month := 0
	fmt.Print("Введите три числа через пробел: исходную сумму, процентную ставку и срок вклада в месяцах\n")
	fmt.Scan(&startSum, &percent, &month)

	percent = (percent / 100) / 12
	finalSum := startSum

	for i := 1; i <= month; i++ {
		finalSum += (finalSum * percent)
	}

	rub := int(finalSum)
	kop := int((finalSum-float64(rub))*100 + 0.5)

	lastDigit := rub % 10
	var strAns string

	if rub%100 >= 11 && rub%100 <= 19 {
		strAns = "рублей"
	} else if lastDigit == 1 {
		strAns = "рубль"
	} else if lastDigit != 0 && lastDigit < 5 {
		strAns = "рубля"
	} else {
		strAns = "рублей"
	}

	fmt.Printf("В конце срока размер вклада составит %d %s %d коп.\n", rub, strAns, kop)
}
