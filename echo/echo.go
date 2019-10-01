// Echo prints its args.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, val := range os.Args {
		fmt.Println(idx, val)
	}
}
