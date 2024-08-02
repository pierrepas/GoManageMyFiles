package main

import (
	"bufio"
	"fmt"
	"go_manage_my_files/pkg/check_duplicate_files"
	"os"
)

func main() {
	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Check for duplicates")
		fmt.Println("2. Exit")

		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')

		switch choice {
		case "1\n":
			check_duplicate_files.CheckForDuplicates()
		case "2\n":
			fmt.Println("Exiting program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
