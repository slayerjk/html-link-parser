package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	// internal packages
	"github.com/slayerjk/html-link-parser/internal/htmlparser"
)

func main() {
	// setting work dir
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("FAILURE: get executable path")
	}
	workDir := path.Dir(exePath)

	// html examples dir
	htmlExamplesDir := workDir + "/html-examples"
	defaultExample := htmlExamplesDir + "/ex1.html"

	// set flags
	htmlFilePath := flag.String("f", defaultExample, "HTML file to parse")
	flag.Parse()

	// parse file
	linkList, err := htmlparser.GetAllHrefsFromFile(*htmlFilePath)
	if err != nil {
		log.Fatalf("FAILURE: get list of link from HTML file(%s):\n\t%v", *htmlFilePath, err)
	}
	fmt.Println(linkList)
}
