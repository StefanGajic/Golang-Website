package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Hello struct {
	Name string
	Time string
}

func main() {

	hello := Hello{"there", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/hello.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			hello.Name = name
		}

		if err := templates.ExecuteTemplate(w, "hello.html", hello); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server starting on 8080")

	fmt.Println(http.ListenAndServe(":8080", nil))
}
