package main

import (
	"fmt"
	"time"
)

type task struct {
	header      string
	description string
	createdAt   string
	status      string
}

func main() {

	toDoList := []task{}
	var userChoose int
	var task task

	fmt.Println("This is to-do-list application.")

	for {
		fmt.Println("Choose function:")
		fmt.Println("1 - Add new task\n2 - Update task\n3 - Delete task\n4 - Task list\n5 - Exit")
		fmt.Scan(&userChoose)

		if userChoose == 1 {
			task = addTask()
			toDoList = append(toDoList, task)
		}

		if userChoose == 2 {
			var n int
			fmt.Println("Enter number of task you want to change")
			fmt.Scan(&n)
			task = addTask()
			toDoList[n-1] = task
		}

		if userChoose == 3 {
			var n int
			fmt.Println("Enter number of task you want to delete")
			fmt.Scan(&n)
			toDoList = append(toDoList[:n-1], toDoList[n:]...)
		}

		if userChoose == 4 {
			taskList(toDoList)
		}

		if userChoose == 5 {
			break
		}

		time.Sleep(1 * time.Second)

	}

}

func addTask() task {
	var newTask task

	fmt.Println("Enter task name")
	fmt.Scan(&newTask.header)

	fmt.Println("Enter task description")
	fmt.Scan(&newTask.description)

	fmt.Println("Enter task date")
	fmt.Scan(&newTask.createdAt)

	fmt.Println("Enter task status")
	fmt.Scan(&newTask.status)

	fmt.Println("New task created")

	return newTask
}

func taskList(tdl []task) {
	for _, val := range tdl {
		fmt.Printf("%v\n", val)
	}
}
