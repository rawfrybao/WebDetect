package main

import (
	"fmt"
	"os"
	"webdetect/internal/detect"
	"webdetect/internal/logger"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("xpathfetch", "Fetch content from a URL using an XPath expression")
	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "URL to fetch content from"})
	xpath := parser.String("x", "xpath", &argparse.Options{Required: true, Help: "XPath expression to extract content"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}

	content := detect.GetContent(*url, *xpath)

	fmt.Println(content)
	logger.Log(content)
}
