package main

import (
	"fmt"
	"net/http"
)

// Student представляє структуру учня
type Student struct {
	ID    int
	Name  string
	Grade string
}

var students = []Student{
	{ID: 1, Name: "Student1", Grade: "5"},
	{ID: 2, Name: "Student2", Grade: "5"},
	{ID: 3, Name: "Student3", Grade: "5"},
	{ID: 4, Name: "Student4", Grade: "6"},
	{ID: 5, Name: "Student5", Grade: "6"},
	{ID: 6, Name: "Student6", Grade: "6"},
}

func main() {
	http.HandleFunc("/student/", studentHandler)
	http.HandleFunc("/grade/", grade)
	http.ListenAndServe(":8080", nil)
}

func grade(w http.ResponseWriter, r *http.Request) {
	grade := r.URL.Path[len("/grade/"):]
	amountOfStudents := 0
	for _, student := range students {
		if student.Grade == grade {
			amountOfStudents++
		}
	}
	if amountOfStudents > 0 {
		fmt.Fprintf(w, "School has : %d, of students in %s grade\n", amountOfStudents, grade)
	} else {
		http.Error(w, "Schood doesn't have that grade currently", http.StatusNotFound)
	}

}
func studentHandler(w http.ResponseWriter, r *http.Request) {
	teacher := r.URL.Query().Get("teacher")
	if teacher != "true" {
		http.Error(w, "You need to be a teacher to have access to students details ", http.StatusUnauthorized)
		return
	}

	id := r.URL.Path[len("/student/"):]
	for _, student := range students {
		if id == fmt.Sprintf("%d", student.ID) {
			fmt.Fprintf(w, "ID: %d, Ім'я: %s, Клас: %s\n", student.ID, student.Name, student.Grade)
			return
		}
	}

	http.Error(w, "We don't have any information about that student", http.StatusNotFound)
}
