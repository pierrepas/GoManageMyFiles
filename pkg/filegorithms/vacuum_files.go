package filegorithms

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func VacuumFiles(inputFile string, destinationFolder string, skipPerLinebreak int) {
	// Reads the input file that should contain a list of files with empty lines.
	// Moves every file listed in the input file to the "destinationFolder", skipping every first "skipPerLinebreak" each new line.

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		log.Printf("Error opening input file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Initialize counters
	lineCount := 0
	skippedCount := 0
	movedFiles := 0

	// Ensure destination folder exists
	if err := os.MkdirAll(destinationFolder, os.ModePerm); err != nil {
		log.Printf("Error creating destination folder: %v\n", err)
		return
	}

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line is empty (linebreak)
		if line == "" {
			lineCount = 0
			skippedCount = 0
			continue
		}

		// Skip files based on skipPerLinebreak
		if skippedCount < skipPerLinebreak {
			skippedCount++
			continue
		}

		// Move the file
		sourceFile := line
		destFile := filepath.Join(destinationFolder, filepath.Base(sourceFile))
		movedFiles++

		err := os.Rename(sourceFile, destFile)
		if err != nil {
			log.Printf("Error moving file %s: %v\n", sourceFile, err)
		} else {
			log.Printf("Moved file: %s to %s\n", sourceFile, destFile)
		}

		lineCount++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
	}
	strMovFil := fmt.Sprint(movedFiles)
	log.Println("Moved " + strMovFil + " files.")
}
