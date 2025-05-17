package main

import (
	"fmt"
	"time"
	task "to-do-list/Tasks"
)

func main() {

	toDoList := []task.Task{}
	var userChoose int

	fmt.Println("This is to-do-list application.")

	for {
		fmt.Println("Choose function:")
		fmt.Println("1 - Add new task\n2 - Update task\n3 - Delete task\n4 - Task list\n5 - Exit")
		fmt.Scan(&userChoose)

		if userChoose == 1 {
			task.AddTask(toDoList)
		}

		if userChoose == 2 {
			task.UpdateTask(toDoList)
		}

		if userChoose == 3 {
			task.DeleteTask(toDoList)
		}

		if userChoose == 4 {
			task.TaskList(toDoList)
		}

		if userChoose == 6 {
			fmt.Printf("Всего задач: %d\n", len(toDoList))
		}

		if userChoose == 5 {
			break
		}

		time.Sleep(1 * time.Second)

	}

}
