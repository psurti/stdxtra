package main

import (
	"fmt"
	"os"
)

func main() {
	// Read the README.md file
	content, err := os.ReadFile("README.md")
	if err != nil {
		fmt.Printf("Error reading README.md: %v\n", err)
		return
	}

	// Output the content
	fmt.Println(string(content))
}
