// Fetch prints the content found at a URL.
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
	defer trace("fetch")()
	const prefix = "http://"
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, prefix) == false {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		fmt.Printf("\n%d bytes read. Http status: %s\n", b, resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func trace(msg string) func() {
	start := time.Now()
	return func() { fmt.Println("Operation", msg, "TimeTaken", time.Since(start)) }
}
