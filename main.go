package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if at least a command was provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: minicli <command> [file]")
		return
	}

	// Get the command (first argument after program name)
	cmd := os.Args[1]
	// Get the filename if provided (second argument after program name)
	file := ""
	if len(os.Args) > 2 {
		file = os.Args[2]
	}

	// Route to the appropriate command function
	if cmd == "cat" {
		cat(file)
	} else if cmd == "ls" {
		ls()
	} else if cmd == "head" {
		head(file)
	} else if cmd == "tail" {
		tail(file)
	} else {
		fmt.Println("Unknown command")
	}
}

// cat command - displays entire file content
func cat(file string) {
	// If no file specified, read from standard input (stdin)
	if file == "" {
		// Copy everything from stdin to stdout
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	// Read the entire file into memory
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Convert byte slice to string and print
	fmt.Print(string(data))
}

// ls command - lists files in current directory
func ls() {
	// Read all directory entries from current directory "."
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print each file/directory name
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

// head command - displays first 10 lines of a file
func head(file string) {
	// Read the entire file
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Split file content into individual lines
	lines := splitLines(string(data))

	// Determine how many lines to show (max 10)
	n := 10
	if len(lines) < n {
		n = len(lines)
	}

	// Print the first n lines
	for i := 0; i < n; i++ {
		fmt.Println(lines[i])
	}
}

// tail command - displays last 10 lines of a file
func tail(file string) {
	// Read the entire file
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Split file content into individual lines
	lines := splitLines(string(data))

	// Determine how many lines to show (max 10)
	n := 10
	if len(lines) < n {
		n = len(lines)
	}

	// Calculate starting index for last n lines
	start := len(lines) - n
	// Print the last n lines
	for i := start; i < len(lines); i++ {
		fmt.Println(lines[i])
	}
}

// splitLines splits a string into lines based on newline characters
func splitLines(s string) []string {
	var lines []string // Slice to store the lines
	var line []byte    // Temporary buffer for current line

	// Loop through each character in the string
	for i := 0; i < len(s); i++ {
		// If we find a newline character, save the current line
		if s[i] == '\n' {
			lines = append(lines, string(line))
			line = []byte{} // Reset line buffer for next line
		} else {
			// Add character to current line buffer
			line = append(line, s[i])
		}
	}

	// If there's any remaining content after last newline, add it as a line
	if len(line) > 0 {
		lines = append(lines, string(line))
	}

	return lines
}
