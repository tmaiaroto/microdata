// Copyright 2015 Lars Wiegman

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/namsral/microdata"
)

func main() {
	var result microdata.Result
	var err error

	baseURL := flag.String("base-url", "http://example.com", "base url to use for the data in the stdin stream.")
	contentType := flag.String("content-type", "utf-8", "content type of the data in the stdin stream.")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s [options] [url]:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nExtract the HTML Microdata from a HTML5 document.")
		fmt.Fprint(os.Stderr, " Provide an URL to a valid HTML5 document or stream a valid HTML5 document through stdin.\n")
	}

	flag.Parse()

	// Items from URL
	if args := flag.Args(); len(args) > 0 {
		urlStr := args[0]
		result, err = microdata.ParseURL(urlStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		printResult(os.Stdout, result)
		return
	}

	// Items from STDIN
	r := os.Stdin
	result, err = microdata.Parse(r, *baseURL, *contentType)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	printResult(os.Stdout, result)
}

// printResult pretty formats and prints the given items in a JSON object.
func printResult(w io.Writer, result microdata.Result) {
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	w.Write(b)
	w.Write([]byte{10})
}