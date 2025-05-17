package task

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	id          int
	header      string
	description string
	createdAt   string
	status      string
}

func AddTask(tdl []Task) {
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

	tdl = append(tdl, newTask)
	newTask.id = len(tdl)
	tdl[newTask.id-1].id = newTask.id

}

func UpdateTask(tdl []Task) {
	var i int
	task := []Task{}
	fmt.Println("Enter number of task you want to change")
	fmt.Scan(&i)
	AddTask(task)
	oldId := tdl[i-1].id
	tdl[i-1] = task[0]
	tdl[i-1].id = oldId
}

func DeleteTask(tdl []Task) []Task {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number of task you want to delete")
	scanner.Scan()
	index := scanner.Text()

	taskID, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println("Invalid number of task")
		return tdl
	}

	for i, t := range tdl {
		if t.id == taskID {
			return append(tdl[:i], tdl[i+1:]...)
		}
	}
	fmt.Println("Task not found")
	return tdl
}

func TaskList(tdl []Task) {
	for _, val := range tdl {
		fmt.Printf("%d. %s %s %s %s\n",
			val.id,
			val.header,
			val.description,
			val.createdAt,
			val.status,
		)
	}
}
