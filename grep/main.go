package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dir := "./"
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			fmt.Println("file", file.Name())
		}
	}
	for _, file := range files {
		if !file.IsDir() {
			fmt.Println("directory", file.Name())
		}
	}
}
