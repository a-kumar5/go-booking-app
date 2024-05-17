package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type todo struct {
	UserID    int    `json:userID`
	ID        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

var form = `
<h1> Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"

	res, err := http.Get(base + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	var item todo

	err = json.Unmarshal(body, &item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	tmpl := template.New("mine")
	tmpl.Parse(form)
	tmpl.Execute(w, item)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
