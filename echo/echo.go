// Echo prints its args.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer trace("echo")()
	for idx, val := range os.Args[1:] {
		fmt.Println(idx, val)
	}
	// fmt.Println(strings.Join(os.Args[1:], " "))
}

func trace(msg string) func() {
	start := time.Now()
	return func() { fmt.Println("Operation", msg, "TimeTaken", time.Since(start)) }
}
