package main

import (
	"flag"
	"fmt"
	"go_manage_my_files/pkg/check_duplicate_files"
)

func main() {
	var actionFlag = flag.String("action", "check_duplicates", "What action the program should take")
	var outputFileFlag = flag.String("output", "duplicates_found.txt", "What file the duplicates should be listed into")
	var pathtoSearchFlag = flag.String("path", ".", "What path the duplicates should be searched from")
	flag.Parse()

	if *actionFlag == "check_duplicates" {
		check_duplicate_files.CheckForDuplicates(*outputFileFlag, *pathtoSearchFlag)
	} else {
		fmt.Println("Flag not recognised")
	}
}
