package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "exercise1.4: %v\n", err)
				continue
			}
			countLines(f, counts, names)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%s\t%d\t%s\n", names[line], n, line)
		}
	}
}

func in(s []string, item string) bool {
	for _, v := range s {
		if v == item {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, names map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !in(names[input.Text()], f.Name()) {
			names[input.Text()] = append(names[input.Text()], f.Name())
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
