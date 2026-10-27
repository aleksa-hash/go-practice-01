package main

import "fmt"

func main() {
	var totalMinutes int
	fmt.Print("Введите количество минут: ")

	// Этот блок ввода дан готовым; вычисления добавьте самостоятельно.
	_, err := fmt.Scan(&totalMinutes)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число.")
		return
	}
	if totalMinutes < 0 {
		fmt.Println("Ошибка: количество минут не должно быть отрицательным.")
		return
	}

	const Hour = 60
	var timehour, timemin int
	timehour = totalMinutes / 60
	timemin = totalMinutes % 60

	fmt.Printf("Полных часов: %d.\n", timehour)
	fmt.Printf("Оставшихся минут: %d.\n", timemin)
}
