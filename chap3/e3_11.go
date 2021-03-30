/*  Enhance comma so that it deals correctly
 *  with floating-point numbers and an optional sign
 */
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, str := range os.Args[1:] {
		fmt.Println(comma(str))
	}
}

func comma(s string) string {

}
