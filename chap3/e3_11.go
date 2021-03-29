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
		fmt.Println(decimal(str))
	}
}

func decimal(s string) string {

	sl := strings.Split(s, ".")

	var sign string

	if strings.HasPrefix(sl[0], "-") || strings.HasPrefix(sl[0], "+") {
		sign = string(sl[0][0])
	}

	if len(sl) == 1 {
		if sign == "" {
			return comma(s)
		} else {
			return sign + comma(s[1:])
		}
	} else {
		if sign == "" {
			return comma(sl[0]) + "." + sl[1]
		} else {
			return sign + comma(sl[0][1:]) + "." + sl[1]
		}
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
		buf.WriteString("," + s[i:i+3])
	}

	return buf.String()
}
