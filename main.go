package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: mytool <command>")
		fmt.Println("Commands: cat, ls, head, tail")
		return
	}

	command := os.Args[1]

	switch command {
	case "cat":
		catCommand()
	case "ls":
		lsCommand()
	case "head":
		headCommand()
	case "tail":
		tailCommand()
	default:
		fmt.Println("Unknown command:", command)
	}
}

func catCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mytool cat <filename>")
		return
	}

	filename := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func lsCommand() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func headCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mytool head <filename>")
		return
	}

	filename := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() && lineCount < 10 {
		fmt.Println(scanner.Text())
		lineCount++
	}
}

func tailCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: mytool tail <filename>")
		return
	}

	filename := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	start := len(lines) - 10
	if start < 0 {
		start = 0
	}
	
	for i := start; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}