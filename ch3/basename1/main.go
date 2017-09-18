// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 72.

// Basename1 reads file names from stdin and prints the base name of each one.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	//	file, err := os.Open("test.txt")
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error found 1: %v", err)
	} else {
		input := bufio.NewReader(file)

		for {
			text, err := input.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil && err != io.EOF {
				panic(err)
			}
			fmt.Println(basename(text))
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+
// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

//!-
