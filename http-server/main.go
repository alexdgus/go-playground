package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	fmt.Println(req.Form)

	type Name struct {
		FirstName string
		LastName  string
	}
	var name Name

	fmt.Println(req.Form)
	if len(req.Form["fn"]) > 0 {
		name.FirstName = req.Form["fn"][0]
	}
	if len(req.Form["ln"]) > 0 {
		name.LastName = req.Form["ln"][0]
	}

	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Println(err)
	}
	err = t.ExecuteTemplate(w, "index.gohtml", name)
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
