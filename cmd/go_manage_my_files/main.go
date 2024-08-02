package main

import (
	"flag"
	"fmt"
	"go_manage_my_files/pkg/filegorithms"
	"log"
	"time"
)

func main() {
	startTime := time.Now()
	log.SetFlags(0) // Removes time stamp from log
	actionFlag := flag.String("action", "check_duplicates", "What action the program should take")
	outputFileFlag := flag.String("output", "duplicates_found.txt", "What file the duplicates should be listed into")
	pathtoSearchFlag := flag.String("path", ".", "What path the duplicates should be searched from")
	flag.Parse()

	if *actionFlag == "check_duplicates" {
		filegorithms.CheckForDuplicates(*outputFileFlag, *pathtoSearchFlag)
	} else {
		fmt.Println("Flag not recognised")
	}
	timeSinceStart := time.Since(startTime)
	log.Println()
	log.Printf("Script took %s", timeSinceStart)
	log.Println()
}
