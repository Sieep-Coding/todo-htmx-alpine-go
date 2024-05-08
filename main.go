package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"isCompleted"`
}

var todos = []Todo{
	{Id: 1, Name: "Learn Go", IsCompleted: false},
	{Id: 2, Name: "Learn Alpine", IsCompleted: false},
	{Id: 3, Name: "Go to the gym", IsCompleted: true},
}

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["index.html"] = template.Must(template.ParseFiles("index.html"))
}

// handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(todos)

	if err != nil {
		log.Fatal(err)
	}

	tmpl := templates["index.html"]
	tmpl.ExecuteTemplate(w, "index.html", map[string]template.JS{"Todos": template.JS(json)})
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
