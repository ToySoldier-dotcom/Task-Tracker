package task

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID          int
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

	fmt.Print("Name - ")
	scanner.Scan()
	newTask.header = scanner.Text()

	fmt.Print("Description - ")
	scanner.Scan()
	newTask.description = scanner.Text()

	ChangeDate(&newTask, scanner)

	ChangeStatus(&newTask, scanner)

	fmt.Println("New task created")

	*tdl = append(*tdl, newTask)
}

func ChangeDate(newtask *Task, scanner *bufio.Scanner) {
	var d, m, y string

	fmt.Println("Enter task date")

	fmt.Print("Day - ")
	scanner.Scan()
	d = scanner.Text()

	fmt.Print("Month - ")
	scanner.Scan()
	m = scanner.Text()

	fmt.Print("Year - ")
	scanner.Scan()
	y = scanner.Text()

	newtask.createdAt = d + "." + m + "." + y + "."
}

func ChangeStatus(newtask *Task, scanner *bufio.Scanner) {
	fmt.Println("Select Status:\n1 - New Task\n2 - Task in Progress\n3 - Finished Task")

	scanner.Scan()
	choose := scanner.Text()

	for {
		if choose == "1" {
			newtask.status = "New Task"
			break
		} else if choose == "2" {
			newtask.status = "Task in Progress"
			break
		} else if choose == "3" {
			newtask.status = "Finished Task"
			break
		} else {
			fmt.Println("You chosed wrong number, try again")
			fmt.Println("Select Status:\n1 - New Task\n2 - Task in Progress\n3 - Finished Task")
			scanner.Scan()
			choose = scanner.Text()
		}
	}
}

func UpdateTask(tdl []Task) ([]Task, int) {
	var i int
	task := []Task{}

	for {
		fmt.Println("Enter number of task you want to change")
		fmt.Scan(&i)
		if i >= 0 && i <= len(tdl) {
			break
		} else {
			fmt.Println("This task does not exist. Try again")
		}
	}

	AddTask(&task)
	return task, i
}

func TaskList(tdl []Task) {
	for _, val := range tdl {
		fmt.Println("--------------------------------------------------")
		fmt.Printf("%d. %s\nDescription - %s\nDate - %s\nStatus - %s\n",
			val.ID,
			val.header,
			val.description,
			val.createdAt,
			val.status,
		)
		fmt.Println("--------------------------------------------------")
	}
}
