package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	noNewline = flag.Bool("n", false, "Do not output the trailing newline")
	logEcho   = flag.Bool("l", false, "Log the echoed message along with timestamp")
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		return
	}

	output := strings.Join(args, " ")

	if *logEcho {
		logEchoMessage(output)
		return
	}

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

func logEchoMessage(message string) {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 Mon 15:04:05.000Z -0700")
	fmt.Printf("%s: %s\r\n", formattedTime, message)
}
