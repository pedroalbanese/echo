package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var noNewline = flag.Bool("n", false, "Do not output the trailing newline")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	output := strings.Join(args, " ")

	if *noNewline {
		os.Stdout.WriteString(output)
	} else {
		newline := "\n"
		if os.PathSeparator == '\\' {
			newline = "\r\n"
		}
		fmt.Print(output + newline)
	}
}
