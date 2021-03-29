/* Write a non-recursive version of comma,  
 * using bytes.Buffer instead of string concatenation
 */
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	for _, str := range os.Args[1:] {
		fmt.Println(comma(str))
	}
}

func comma(s string) string {
	var buf bytes.Buffer
	
	n := len(s)

	if n <= 3 {
		return s
	}

	i := n % 3
	
	if i == 0 {
		buf.WriteString(s[0:3])
		i = 3
	} else {
		buf.WriteString(s[:i])
	}

	
	for ; i < n; i = i + 3 {
		buf.WriteString( "," + s[i:i+3])
	}

	return buf.String() 
}