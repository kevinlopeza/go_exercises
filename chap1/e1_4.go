//This program finds duplicate lines in
//different files
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//This map will hold the count of how many times a line
	//appears within a given file
	counts_per_file := make(map[string]map[string]int)

	files := os.Args[1:]

	if len(files) != 0 {
		for _, file := range files {
			byteSlice, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			lines := strings.Split(string(byteSlice), "\n")
			for _, line := range lines {
				if counts_per_file[line] == nil {
					counts_per_file[line] = make(map[string]int)
				}
				counts_per_file[line][file]++
			}
		}

		//We now print the results
		for line, submap := range counts_per_file {
			for filename, n := range submap {
				if n > 1 {
					fmt.Printf("%s\t%d\t%s\n", filename, n, line)
				}
			}
		}
	}
}
