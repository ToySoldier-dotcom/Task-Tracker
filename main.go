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

	fmt.Println("This is to-do-list application.")

	for {
		fmt.Println("Choose function:")
		fmt.Println("1 - Add new task\n2 - Update task\n3 - Delete task\n4 - Task list\n5 - Exit")
		fmt.Scan(&userChoose)

		if userChoose == 1 {
			addTask(&toDoList)
		}

		if userChoose == 2 {
			updateTask(toDoList)
		}

		if userChoose == 3 {
			deleteTask(toDoList)
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

func addTask(tdl *[]task) {
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

	*tdl = append(*tdl, newTask)
}

func updateTask(tdl []task) {
	var i int
	task := []task{}
	fmt.Println("Enter number of task you want to change")
	fmt.Scan(&i)
	addTask(&task)
	tdl[i-1] = task[0]
}

func deleteTask(tdl []task) {
	var i int
	fmt.Println("Enter number of task you want to delete")
	fmt.Scan(&i)
	tdl = append(tdl[:i-1], tdl[i:]...)
}

func taskList(tdl []task) {
	for _, val := range tdl {
		fmt.Printf("%v\n", val)
	}
}
