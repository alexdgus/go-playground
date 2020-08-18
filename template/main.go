package main

import (
	"html/template"
	"log"
	"os"
)

type team struct {
	City string
	Name string
	Seed int8
}

func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	lakers := team{"Los Angeles", "Lakers", 1}
	t.Execute(os.Stdout, lakers)
}
