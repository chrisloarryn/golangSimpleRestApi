package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"log"
)

type task struct {
	ID      int    `json:ID`
	Name    string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some content",
	},
	{
		ID:      2,
		Name:    "Task Two",
		Content: "Some content",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please insert a valid task!")
	}
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	// data type / StatusCreated to give feedback / create task
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my REST API jj")
}

func main() {
	// fmt.Println("Hello from GoLang")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", router))
}
