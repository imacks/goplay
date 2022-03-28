package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	goplayapi "github.com/haya14busa/goplay"
)

var (
	runCode bool
	shareCode bool
	openInBrowser bool
	apiEndpoint string
)

func init() {
	flag.Usage = func() {
		fmt.Println("Usage: goplay [OPTION]... [FILE] (reads from stdin if FILE is - or omitted)")
		fmt.Println("")
		flag.PrintDefaults()
	}

	flag.BoolVar(&runCode, "run", true, "compile and run Go program on The Go Playground")
	flag.BoolVar(&shareCode, "share", false, "share code on the Go Playground")
	flag.BoolVar(&openInBrowser, "openbrowser", false, "open in browser automatically")
	flag.StringVar(&apiEndpoint, "endpoint", "https://play.golang.org", "Goplay server endpoint")
}

func main() {
	flag.Parse()

	if !runCode && !shareCode {
		os.Exit(0)
	}

	var file *os.File
	if len(flag.Args()) > 0 {
		filePath := flag.Arg(0)
		if filePath != "-" {
			// use stdin
			var err error
			file, err = os.Open(filePath)
			fatal(err)
			defer file.Close()
		}
	} else {
		file = os.Stdin
	}

	codeBytes, err := io.ReadAll(file)
	fatal(err)
	if string(codeBytes) == "" {
		fatal(fmt.Errorf("INPUT is empty"))
	}

	client := goplayapi.Client{BaseURL: apiEndpoint}
	if runCode {
		if err := client.Run(bytes.NewReader(codeBytes), os.Stdout, os.Stderr); err != nil {
			fatal(err)
		}
	}

	if shareCode {
		url, err := client.Share(bytes.NewReader(codeBytes))
		fatal(err)

		fmt.Println(url)
		if openInBrowser {
			fatal(invokeBrowser(url).Start())
		}
	}
}

func fatal(err error) {
	if err == nil {
		return
	}

	fmt.Printf("fatal: %v\n", err)
	os.Exit(1)
}