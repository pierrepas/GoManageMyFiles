package filegorithms

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func WriteHashMap(outputFile string, pathToSearch string) error {
	// Open the output file in append mode
	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	prefixString := "sha256:"
	hashedFiles := 0
	if err != nil {
		return fmt.Errorf("error opening output file: %v", err)
	}
	defer file.Close()

	// Create a buffered writer for efficient writing
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Create a wait group to synchronize goroutines
	var wg sync.WaitGroup

	// Create a mutex to ensure thread-safe writing to the file
	var mu sync.Mutex

	// Create a channel to limit the number of concurrent goroutines
	semaphore := make(chan struct{}, 100) // Adjust this number based on your system's capabilities

	// Function to process each file
	processFile := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error accessing path %s: %v", path, err)
			return nil // Continue walking despite errors
		}

		if !info.IsDir() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				semaphore <- struct{}{}        // Acquire semaphore
				defer func() { <-semaphore }() // Release semaphore

				// Calculate the relative path
				relPath, err := filepath.Rel(pathToSearch, path)
				if err != nil {
					log.Printf("Error getting relative path for %s: %v", path, err)
					return
				}

				// Calculate the file hash
				hash, err := calculateFileHash(path)
				if err != nil {
					log.Printf("Error calculating hash for %s: %v", path, err)
					return
				}

				// Write to the file in a thread-safe manner
				mu.Lock()
				_, err = fmt.Fprintf(writer, "%s\n%s%s\n\n", relPath, prefixString, hash)
				mu.Unlock()
				hashedFiles++
				if err != nil {
					log.Printf("Error writing to output file: %v", err)
				}
			}()
		}
		return nil
	}

	// Walk the directory
	err = filepath.Walk(pathToSearch, processFile)
	if err != nil {
		return fmt.Errorf("error walking the path %s: %v", pathToSearch, err)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	strHashFiles := fmt.Sprint(hashedFiles)
	log.Println("Hashed " + strHashFiles + " files.")

	return nil
}

// Helper function to calculate the SHA256 hash of a file
func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
