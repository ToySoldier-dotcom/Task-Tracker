package task

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	header      string
	description string
	createdAt   string
	status      string
}

func AddTask(tdl *[]Task) {
	var newTask Task
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		break
	}

	fmt.Println("Enter task name")
	scanner.Scan()
	newTask.header = scanner.Text()

	fmt.Println("Enter task description")
	scanner.Scan()
	newTask.description = scanner.Text()

	fmt.Println("Enter task date")
	scanner.Scan()
	newTask.createdAt = scanner.Text()

	fmt.Println("Enter task status")
	scanner.Scan()
	newTask.status = scanner.Text()

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
