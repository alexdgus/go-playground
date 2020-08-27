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

	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Println(err)
	}
	err = t.ExecuteTemplate(w, "index.gohtml", req)
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
