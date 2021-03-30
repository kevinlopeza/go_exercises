/* Write a function that reports whether two strings
 * are anagrams of each other. 
 */
package main 

import (
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println(anagrams(os.Args[1], os.Args[2]))
}

func anagrams(s1, s2 string) bool {

	if len(s1) != len(s2) { return false }

	b1, b2 := []byte(s1), []byte(s2)

	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

	if bytes.Compare(b1, b2) == 0 {
		return true
	} else {
		return false
	}
}