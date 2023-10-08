package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	ID   int
	Name string
	Date time.Time
}

var tasks []Task

func main() {
	tasks = append(tasks, Task{ID: 1, Name: "Task1", Date: time.Date(2023, time.October, 10, 0, 0, 0, 0, time.UTC)})
	tasks = append(tasks, Task{ID: 2, Name: "Task2", Date: time.Date(2023, time.October, 15, 0, 0, 0, 0, time.UTC)})
	tasks = append(tasks, Task{ID: 3, Name: "Task3", Date: time.Date(2023, time.October, 12, 0, 0, 0, 0, time.UTC)})
	tasks = append(tasks, Task{ID: 4, Name: "Task4", Date: time.Date(2023, time.October, 12, 0, 0, 0, 0, time.UTC)})

	http.HandleFunc("/tasks", getTasksHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("use next format to get tasks per each date /tasks?from=2023-10-12 for example")

}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	dateStr := r.URL.Query().Get("from")
	if dateStr == "" {
		http.Error(w, "date is not defined, use like /tasks?from=2023-10-12 for example", http.StatusBadRequest)
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "date format is not correct, use like /tasks?from=2023-10-12 for example ", http.StatusBadRequest)
		return
	}

	TasksLists := []Task{}
	for _, task := range tasks {
		if task.Date.Year() == date.Year() && task.Date.Month() == date.Month() && task.Date.Day() == date.Day() {
			TasksLists = append(TasksLists, task)
		}
	}

	fmt.Fprintf(w, "List of tasks at  %s:\n", dateStr)
	for _, task := range TasksLists {
		fmt.Fprintf(w, "ID: %d, Task: %s \n", task.ID, task.Name)
	}
}
