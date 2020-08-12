package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var recursive bool
	var dirs []string
	for i, arg := range os.Args {
		if i == 0 {
			continue
		}
		if arg == "-r" {
			recursive = true
			fmt.Println(arg)
		} else {
			dirs = append(dirs, arg)
			fmt.Println(arg)
		}
	}

	dirs = append(dirs, "./")

	if recursive {
		for _, dir := range dirs {
			fmt.Println(dir)
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
	}
}
