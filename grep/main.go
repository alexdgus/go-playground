package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	}
}

func findFiles(dirPath string, recursive bool) {
	fmt.Println(dirPath)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() && recursive {
			findFiles(file.Name(), recursive)
		} else {
			parse(file.Name(), "for")
		}
	}
}

func parse(path string, searchStr string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	line := 1
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchStr) {
			fmt.Println("String found:", line)
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		// Handle the error
	}
}
