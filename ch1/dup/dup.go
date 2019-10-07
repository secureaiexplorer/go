// Dup prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	dupfile := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, dupfile)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts, dupfile)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("count = %d\t files = %s\ntext = %s\n", n, dupfile[line], line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, dupfile map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			if strings.Contains(dupfile[input.Text()], f.Name()) == false {
				dupfile[input.Text()] += f.Name()
				dupfile[input.Text()] += " "
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}
