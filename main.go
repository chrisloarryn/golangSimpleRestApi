package main

import ("fmt"
		"github.com/gorilla/mux"
		"net/http"
		"log"	
)

type task struct {
	ID int `json:ID`
	Name string `json:Name`
	Content string `json:Content`
}

type allTasks []task

var tasks = allTasks {
	{
		ID: 1,
		Name: "Task One",
		Content: "Some content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my API")
}

func main() {
	// fmt.Println("Hello from GoLang")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)

	log.Fatal(http.ListenAndServe(":3000", router))
}