package main

import (
	"fmt"
	"time"
)

func main() {
	twelveHourTime := "03:04:00 PM"

	convertTime := convert(twelveHourTime)
	fmt.Println(convertTime)
}
func convert(twelveHourTime string) (twentyFourHourTime string) {

	twelveHourFormat := "03:04:00 PM"
	timeObj, err := time.Parse(twelveHourFormat, twelveHourTime)
	if err != nil {
		fmt.Println("Некоректний формат часу")
		return
	}
	// Конвертуємо в двадцятичотирьохгодинний час
	twentyFourHourFormat := "15:04:00"
	twentyFourHourTime = timeObj.Format(twentyFourHourFormat)

	// Виводимо результат
	return twentyFourHourTime
}
