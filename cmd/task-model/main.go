package main

import (
	"fmt"
)

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

// Возвращает "#ID Название [статус], приоритет N"
func (t Task) Summary() string {
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

// Присваевает статус done исходной задачи
func (t *Task) Complete() {
	t.Status = "done"
}

// Проверяет диапозон 1-3
func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
}

// При совпадении ID возвращает &tasks[i], а иначе nil
func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	return nil
}

// Возвращает новый срез задач с заданным статусом , а иначе nil
func tasksByStatus(tasks []Task, status string) []Task {
	var rez []Task
	for i := range tasks {
		if tasks[i].Status == status {
			rez = append(rez, tasks[i])
		}
	}
	return rez
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}

	found := findTaskByID(tasks, 2) //Нашли ID = 2
	if found != nil {
		found.Complete()     //Завершили задачу с ID = 2
		found.SetPriority(3) //Установили приоритет 3
	}
	count := 0
	for i := range tasks {
		if tasks[i].Status == "done" {
			count++ //Посчитали завершённые задачи
		}
	}
	for _, task := range tasks {
		fmt.Println(task.Summary())
	}
	fmt.Printf("Задач завершено: %d\n", count)

	if findTaskByID(tasks, 99) == nil { //проверка случая когда нет совпадений ID
		fmt.Println("Задача не найдена\n")
	}
	if tasksByStatus(tasks, "none") == nil { //проверка случая когда нет совпадений статуса
		fmt.Println("Задача с таким статусом не найдена\n")
	}
	dones := tasksByStatus(tasks, "done")
	if dones != nil {
		fmt.Println("Задач с таким статусом:\n", len(dones))
	}
	for _, task := range dones {
		fmt.Println(task.Summary()) //Вывод задач с заданым статусом
	}
}
