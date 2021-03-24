//this program fetches URLs in parallel and
//places the contents received in files
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	secs := time.Since(start).Seconds()
	fmt.Printf("main took %.2fs\n", secs)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	filename := strings.Split(url, ".")[1] + ".html"

	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	bytesTransferred, err := io.Copy(file, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading: %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("Succesfully wrote %d bytes in %s. The operation took %.2fs.", bytesTransferred, filename, secs)

}
