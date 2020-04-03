// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Result struct {
	count     int
	filenames map[string]bool
}

func main() {
	counts := make(map[string]*Result)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, res := range counts {
		if res.count > 1 {
			fmt.Printf("%d\t%s\n", res.count, line)
			for key, _ := range res.filenames {
				fmt.Printf("-------------- %s\n", key)
			}
		}
	}
}
func countLines(f *os.File, counts map[string]*Result) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		txt := input.Text()
		if _, ok := counts[txt]; !ok {
			counts[txt] = &Result{0, make(map[string]bool)}
		}
		counts[txt].count += 1
		if !counts[txt].filenames[f.Name()] {
			counts[txt].filenames[f.Name()] = true
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
