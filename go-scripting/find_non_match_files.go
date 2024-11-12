package main

import (
	"bufio"
	"fmt"
	"os"
)

// Prints non-matching files given a list of filenames inside a file as input.
// Requires {path of filelist} and {path of directory} to compare to (should be a subset of filelist) as arguments.
func find_non_match_files() {
	list_filepath := os.Args[1] // Args[0] is the program name so we start at index 1
	dirpath := os.Args[2]       // subset of filenames in the list_filepath

	if list_filepath == "" || dirpath == "" { // WARN: this check does not properly work lol
		fmt.Println("No args specified. Need 2: {file-path-of-list} and {directory-path-to-compare-list-to}")
		os.Exit(1)
	}

	file, err := os.Open(list_filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// input file
	var fname_inputfile []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		fname_inputfile = append(fname_inputfile, s.Text())
	}
	if err := s.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fname_fromdir := make(map[string]bool)
	dirfiles, err := os.ReadDir(dirpath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
	}
	for _, entry := range dirfiles {
		fname_fromdir[entry.Name()] = true
	}

	// var non_match_files []string
	// compare fname_inputfile and fname_fromdir
	for _, filename := range fname_inputfile {
		if !fname_fromdir[filename] {
			fmt.Println(filename)
		}
	}
}

func main() {
	find_non_match_files()
}
