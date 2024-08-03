package main

import (
	"flag"
	"go_manage_my_files/pkg/filegorithms"
	"log"
	"time"
)

func main() {
	/*
		-action=[check_duplicates | check_names | vacuum_files] (default:check_duplicates)
			check_duplicates: checks all duplicate files in the "path" parameter directory and writes them into "filename".
			vacuum_files: moves all files listed in the "filename" file into the "path" folder, except for the first one and first after each linebreak.
		-filename=<filename, default="duplicates_found.txt">
		-path=<path, default=".">
	*/
	startTime := time.Now()
	log.SetFlags(0) // Removes time stamp from log
	actionFlag := flag.String("action", "check_duplicates", "Action the program will take")
	fileName := flag.String("filename", "files.txt", "File of the action")
	pathOfAction := flag.String("path", ".", "Path of the action")
	flag.Parse()

	if *actionFlag == "check_duplicates" {
		filegorithms.CheckDuplicateFiles(*fileName, *pathOfAction)
	} else if *actionFlag == "check_names" {
		filegorithms.CheckDuplicateNames(*fileName, *pathOfAction)
	} else if *actionFlag == "vacuum_files" {
		filegorithms.VacuumFiles(*fileName, *pathOfAction, 1)
	} else if *actionFlag == "write_hashmap" {
		filegorithms.WriteHashMap(*fileName, *pathOfAction)
	} else {
		log.Fatalln("Arguments not recognised, please try again")
	}
	timeSinceStart := time.Since(startTime)
	log.Println()
	log.Printf("Script took %s", timeSinceStart)
	log.Println()
}
