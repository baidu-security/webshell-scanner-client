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
		fmt.Printf("WebShell scanner API client - Copyright ©2018 Baidu Inc.\n")
		fmt.Printf("For more details visit: https://scanner.baidu.com\n\n")
		fmt.Printf("Usage: %s /tmp/a.php /tmp/b.php ...\n", os.Args[0])
		return
	}

	for _, arg := range os.Args[1:] {
		scanner.ProcessFile(arg)
	}
}
