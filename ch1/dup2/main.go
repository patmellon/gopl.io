// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type FileData struct {
	name   string
	counts map[string]int
}

func main() {
	totalCounts := make(map[string]*FileData)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, totalCounts, files[0])
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, totalCounts, arg)
			f.Close()
		}
	}

	for _, n := range totalCounts {
		for word, count := range n.counts {
			if count > 1 {
				fmt.Printf("%d\t%s\n", count, word+n.name)
			}
		}

	}
}

func countLines(f *os.File, totalCounts map[string]*FileData, filename string) {
	fileData := new(FileData)
	fileData.name = filename
	fileData.counts = make(map[string]int)

	input := bufio.NewScanner(f)

	for input.Scan() {
		fileData.counts[input.Text()]++
	}

	totalCounts[filename] = fileData
	// NOTE: ignoring potential errors from input.Err()
}

/*
	a := new(FileData)
	a.name = "test"
	a.counts = make(map[string]int)

	a.counts["k2"] = 13

	totalCounts["hello"] = a
*/
//!-
