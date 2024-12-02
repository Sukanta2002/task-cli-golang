package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/rodaine/table"
)

var path = "./task.json"
var tasks []Task

// types
type Task struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func main() {
	file := CreateOrOpenfile()
	GetTasksFromFile()
	defer CloseFile(file)

	args := os.Args

	if len(args) < 1 {
		fmt.Println("Enter some data")
		return
	}

	switch args[1] {
	case "add":
		AddTask(strings.TrimSpace(args[2]))
	case "list":
		if len(args) < 3 {
			ListTask("all")
		} else {
			ListTask(args[2])
		}

	case "delete":
		if len(args) < 3 {
			fmt.Println("Enter the id to delete ")

		} else {
			DeleteTask(args[2])
		}
	case "update":
		if len(args) < 4 {
			fmt.Println("Enter the id to update along with ")
		} else {
			UpdateTask(args[2], args[3])
		}
	case "mark-in-progress":
		fmt.Println("running mark in progerss")
		if len(args) < 3 {
			fmt.Println("Enter the id to mark in progress")
		} else {
			ChangeStatus(args[2], "IN-PROGRESS")
		}
	case "mark-done":
		if len(args) < 3 {
			fmt.Println("Enter the id to mark done")
		} else {
			ChangeStatus(args[2], "DONE")
		}
	}
	defer SaveFile(file)

}

func CreateOrOpenfile() *os.File {
	file, e := os.Open(path)

	if e != nil {
		if os.IsNotExist(e) {
			file, e = os.Create(path)
			if e != nil {
				panic(e)
			}
		} else {
			panic(e)
		}

	}
	return file

}
func CloseFile(file *os.File) {
	file.Close()
}

func AddTask(task string) {
	time := time.Now().Format("2006-01-02T15:04:05-0700")
	id := uuid.New().String()
	data := Task{Description: task, Id: id, Status: "TODO", CreatedAt: time, UpdatedAt: time}

	tasks = append(tasks, data)

}

func SaveFile(file *os.File) {
	data, _ := json.MarshalIndent(tasks, "", "\t")
	os.WriteFile(path, data, os.ModeAppend)
}

func GetTasksFromFile() {
	data, _ := os.ReadFile(path)
	if len(data) == 0 {
		return
	}
	err := json.Unmarshal(data, &tasks)

	if err != nil {
		panic(err)
	}
}

func ListTask(status string) {
	tbl := table.New("ID", "TASK", "STATUS", "CREATED AT", "UPDATED AT")
	switch status {
	case "all":
		for _, v := range tasks {
			tbl.AddRow(v.Id, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
		}
		tbl.Print()
	case "DONE":
		for _, v := range tasks {
			if v.Status == status {
				tbl.AddRow(v.Id, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
			}
		}
		tbl.Print()
	case "TODO":
		for _, v := range tasks {
			if v.Status == status {
				tbl.AddRow(v.Id, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
			}
		}
		tbl.Print()
	case "IN-PROGRESS":
		for _, v := range tasks {
			if v.Status == status {
				tbl.AddRow(v.Id, v.Description, v.Status, v.CreatedAt, v.UpdatedAt)
			}
		}
		tbl.Print()
	}

}

func DeleteTask(id string) {
	index := -1
	for i, v := range tasks {
		if v.Id == id {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("No task with this id")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
}

func UpdateTask(id string, desc string) {
	index := -1
	for i, v := range tasks {
		if v.Id == id {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("No task with this id")
		return
	}
	updateTime := time.Now().Format("2006-01-02T15:04:05-0700")
	tasks[index].Description = desc
	tasks[index].UpdatedAt = updateTime
}

func ChangeStatus(id string, status string) {
	index := -1
	for i, v := range tasks {
		if v.Id == id {
			index = i
		}
	}
	if index == -1 {
		fmt.Println("No task with this id")
		return
	}
	tasks[index].Status = status
}
