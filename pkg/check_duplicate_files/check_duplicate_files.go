package check_duplicate_files

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CheckForDuplicates() {
	fmt.Print("Enter the folder path to check for duplicates: ")
	reader := bufio.NewReader(os.Stdin)
	folderPath, _ := reader.ReadString('\n')
	folderPath = folderPath[:len(folderPath)-1] // Remove newline character

	fileHashes := make(map[string][]string)

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
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

	duplicatesFound := false
	for _, paths := range fileHashes {
		if len(paths) > 1 {
			duplicatesFound = true
			fmt.Println("Duplicate files found:")
			for _, path := range paths {
				fmt.Println(path)
			}
			fmt.Println()
		}
	}

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
