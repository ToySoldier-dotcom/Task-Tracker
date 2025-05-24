package main

import (
	"fmt"
	"time"
	task "to-do-list-mod/Tasks"
)

func main() {

	toDoList := []task.Task{}
	var userChoose int

	fmt.Println("This is to-do-list application.")

	for {
		fmt.Println("Choose function:")
		fmt.Println("1 - Add new task\n2 - Update task\n3 - Task list\n4 - Exit")
		fmt.Scan(&userChoose)

		switch {
		case userChoose == 1:
			task.AddTask(&toDoList)
			id := len(toDoList)
			toDoList[id-1].ID = id
		case userChoose == 2:
			task, i := task.UpdateTask(toDoList)
			oldId := toDoList[i-1].ID
			toDoList[i-1] = task[0]
			toDoList[i-1].ID = oldId
		case userChoose == 3:
			task.TaskList(toDoList)
		case userChoose == 4:
			return
		case userChoose == 5:
			fmt.Printf("Всего задач: %d\n", len(toDoList))
		}

		time.Sleep(1 * time.Second)

	}

}
