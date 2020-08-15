package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func parseInput() (recursive bool, searchString string, dirs []string) {
	r := false
	var ss string
	dirList := []string{}
	for i, arg := range os.Args {
		if i == 0 {
			continue
		} else if i == 1 && arg == "-r" {
			r = true
			fmt.Println(arg)
		} else if i == 1 {
			ss = arg
		} else if i == 2 && recursive {
			ss = arg
		} else {
			dirList = append(dirs, arg)
			fmt.Println(arg)
		}
	}
	if len(dirList) == 0 {
		dirList = append(dirList, ".")
	}
	return r, ss, dirList
}

func main() {
	recursive, searchString, dirs := parseInput()

	for _, input := range os.Args {
		fmt.Println(input)
	}
	fmt.Println("dirs found")
	fmt.Println(dirs)

	var files []string
	for _, dir := range dirs {
		files = findFiles(dir, recursive)
	}
	fmt.Println("files found")
	fmt.Println(files)
	for _, file := range files {
		parse(file, searchString)
	}
}

func findFiles(dirPath string, recursive bool) []string {
	var filePaths []string
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return filePaths
	}

	for _, file := range files {
		if file.IsDir() && recursive {
			filePaths = append(filePaths, findFiles(file.Name(), recursive)...)
		} else if !file.IsDir() {
			filePaths = append(filePaths, dirPath+"/"+file.Name())
		}
	}
	return filePaths
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
			fmt.Println("Line:", line, scanner.Text())
		}
		line++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
