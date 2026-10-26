package main

import "fmt"

func main() {
	const course = "Бэкенд-разработка на языке Go"
	var name string = "Даша"
	completed := 2
	var group string = "Учебная-03"
	interest := true

	fmt.Printf("Участник:%s\n", name)
	fmt.Printf("Группа:%s\n", group)
	fmt.Println("Курс:", course)
	fmt.Printf("Завершено работ:%d\n", completed)
	fmt.Printf("Интерес к бэкенду:%t\n", interest)

}
