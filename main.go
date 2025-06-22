// Author: Vedant Gupta
// Command-line tool: lastused
// Description: Lists recently modified files in the current directory, sorted by time.
// Usage:
//     lastused              # shows all files sorted by recent usage
//     lastused -n 5         # shows top 5 files
//     lastused --ext .go    # only show .go files

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

type fileInfo struct {
	Path    string
	ModTime time.Time
}

func main() {
	// Parse flags
	n := flag.Int("n", 0, "Number of results to show")
	ext := flag.String("ext", "", "Only include files with this extension (e.g. .go)")
	flag.Parse()

	// Get current directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: could not get current directory:", err)
		os.Exit(1)
	}

	var files []fileInfo

	// Walk through the directory and collect file info
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil // skip problematic files
		}
		if !info.IsDir() {
			if *ext != "" && filepath.Ext(path) != *ext {
				return nil // skip unmatched extension
			}
			files = append(files, fileInfo{
				Path:    path,
				ModTime: info.ModTime(),
			})
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error: failed to walk through directory:", err)
		os.Exit(1)
	}

	// Sort files by modification time (descending)
	sort.Slice(files, func(i, j int) bool {
		return files[i].ModTime.After(files[j].ModTime)
	})

	// Limit the number of results
	limit := len(files)
	if *n > 0 && *n < limit {
		limit = *n
	}

	// Display results with precise timestamp
	for i := 0; i < limit; i++ {
		rel, _ := filepath.Rel(dir, files[i].Path)
		formattedTime := files[i].ModTime.Format("2006-01-02 15:04:05")
		fmt.Printf("%-40s — modified on %s\n", rel, formattedTime)
	}
}
