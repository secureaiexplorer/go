// Fetchall fetches URLs in parallel and reports their times and sizes.
// Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by running fetchall
// twice in succession to see whether the reported time changes much. Do you get the same content
// each time? Modify fetchall to print its output to a file so it can be examined.
// Exercise 1.11: Try fetchall with longer argument lists, such as samples from the top million web sites available
// at alexa.com. How does the program behave if a web site just doesn’t respond? (Section 8.9 describes
// mechanisms for coping in such cases.)

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	const prefix = "http://"
	outfile := os.Args[1]
	start := time.Now()
	f, err := os.OpenFile(outfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	ch := make(chan string)
	for _, url := range os.Args[2:] {
		if strings.HasPrefix(url, prefix) == false {
			url = prefix + url
		}
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[2:] {
		fmt.Fprintln(f, <-ch) // receive from channel ch
	}
	fmt.Fprintln(f, time.Since(start).Seconds(), "elapsed")
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
