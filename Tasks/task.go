package task

import "fmt"

type Task struct {
	header      string
	description string
	createdAt   string
	status      string
}

func AddTask(tdl *[]Task) {
	var newTask Task

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

func UpdateTask(tdl []Task) {
	var i int
	task := []Task{}
	fmt.Println("Enter number of task you want to change")
	fmt.Scan(&i)
	AddTask(&task)
	tdl[i-1] = task[0]
}

func DeleteTask(tdl []Task) {
	var i int
	fmt.Println("Enter number of task you want to delete")
	fmt.Scan(&i)
	tdl = append(tdl[:i-1], tdl[i:]...)
}

func TaskList(tdl []Task) {
	for _, val := range tdl {
		fmt.Printf("%v\n", val)
	}
}
