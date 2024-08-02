package check_duplicate_files

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CheckForDuplicates(outputFile string, pathToSearch string) {
	fmt.Println("Searching the path: " + pathToSearch)
	fmt.Println("Writing to file: " + outputFile)
	fmt.Println()
	f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatal(err)
	}

	fileHashes := make(map[string][]string)

	err = filepath.Walk(pathToSearch, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			hash, err := calculateMD5(path)
			if err != nil {
				return err
			}
			fileHashes[hash] = append(fileHashes[hash], path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
		return
	}

	writer := bufio.NewWriter(f)
	duplicatesFound := false
	for _, paths := range fileHashes {
		if len(paths) > 1 {
			duplicatesFound = true
			fmt.Println("Duplicate files found:")
			for _, path := range paths {
				fmt.Println(path)
				writer.WriteString(path + "\n")
			}
			fmt.Println()
			writer.WriteString("\n")
		}
	}

	writer.Flush()

	if !duplicatesFound {
		fmt.Println("No duplicate files found.")
	}
}

func calculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
