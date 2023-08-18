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
	format    = flag.String("f", "2006-01-02 Mon 15:04:05.000Z -0700", "Specify the format for the log timestamp (options: UnixDate, Kitchen, \nRFC3339Nano, RFC3339, RFC1123, RFC1123Z, RubyDate, RFC822, RFC822Z, \nRFC850)")
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

	var newline string
	if os.PathSeparator == '\\' {
		newline = "\r\n"
	} else {
		newline = "\n"
	}

	if *noNewline {
		os.Stdout.WriteString(output)
	} else {
		fmt.Print(output + newline)
	}
}

func logEchoMessage(message string) {
	currentTime := time.Now()
	var formattedTime string

	switch *format {
	case "UnixDate":
		formattedTime = currentTime.Format(time.UnixDate)
	case "Kitchen":
		formattedTime = currentTime.Format(time.Kitchen)
	case "RFC3339Nano":
		formattedTime = currentTime.Format(time.RFC3339Nano)
	case "RFC3339":
		formattedTime = currentTime.Format(time.RFC3339)
	case "RFC1123":
		formattedTime = currentTime.Format(time.RFC1123)
	case "RFC1123Z":
		formattedTime = currentTime.Format(time.RFC1123Z)
	case "RubyDate":
		formattedTime = currentTime.Format(time.RubyDate)
	case "RFC822":
		formattedTime = currentTime.Format(time.RFC822)
	case "RFC822Z":
		formattedTime = currentTime.Format(time.RFC822Z)
	case "RFC850":
		formattedTime = currentTime.Format(time.RFC850)
	default:
		formattedTime = currentTime.Format(*format)
	}

	var newline string
	if os.PathSeparator == '\\' {
		newline = "\r\n"
	} else {
		newline = "\n"
	}

	fmt.Printf("%s %s%s", formattedTime, message, newline)
}
