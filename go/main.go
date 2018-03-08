package main

import (
	"fmt"
	"os"
	// "flag"
	"scanner"
)

var (
// debug = flag.Bool("debug", false, "Enable debug output")
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Webdir scanner API client - Version 2018-03-08\nCopyright Baidu Inc.\n\n")
		fmt.Printf("Usage: %s /tmp/a.php /tmp/b.php ...\n", os.Args[0])
		return
	}

	for _, arg := range os.Args[1:] {
		scanner.ProcessFile(arg)
	}
}
