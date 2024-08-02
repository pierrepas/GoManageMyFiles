package filegorithms

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CheckDuplicateNames(outputFile string, pathToSearch string) {
	// Parses all the files in the pathToSearch directory and prints a list of duplicate names to outputFile.

	log.Println("Checking for duplicate file names.")
	log.Println("Searching the path:", pathToSearch)
	log.Println("Writing to file:", outputFile)
	log.Println()
	duplicateCount := 0

	// Map to store filenames and their paths
	fileMap := make(map[string][]string)

	// Walk through the directory
	err := filepath.Walk(pathToSearch, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			// Get the filename without the path
			filename := filepath.Base(path)
			// Add the path to the slice of paths for this filename
			fileMap[filename] = append(fileMap[filename], path)
		}
		return nil
	})

	if err != nil {
		log.Printf("Error walking through directory: %v\n", err)
		return
	}

	// Open the output file
	file, err := os.Create(outputFile)
	if err != nil {
		log.Printf("Error creating output file: %v\n", err)
		return
	}
	defer file.Close()

	// Check for duplicates and write to file
	for filename, paths := range fileMap {
		if len(paths) > 1 {
			// Write the duplicate filename
			log.Printf("Duplicate filename found: %s\n", filename)
			duplicateCount++
			// Write all paths for this filename
			for _, path := range paths {
				fmt.Fprintf(file, "%s\n", path)
			}
			fmt.Fprintln(file) // Add a blank line for readability
		}
	}
	dupCntStr := fmt.Sprint(duplicateCount)
	log.Println(dupCntStr + " duplicates found.")
	log.Printf("Duplicate filenames have been written to %s\n", outputFile)
}
