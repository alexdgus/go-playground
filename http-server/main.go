package main

import (
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Println(err)
	}
	err = t.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
