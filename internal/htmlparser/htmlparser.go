package htmlparser

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Get raw HTML data from .html file
func getHTMLNodeFromFile(htmlFilePath string) (*html.Node, error) {
	// open HTML file
	htmlFile, err := os.Open(htmlFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open html file(%s):\n\t%v", htmlFilePath, err)
	}

	// get raw HTML
	result, err := html.Parse(htmlFile)
	if err != nil {
		return nil, fmt.Errorf("failed to parse opened html file(%s):\n\t%v", htmlFilePath, err)
	}

	return result, nil
}

// Get all hrefs of <a> of html file
func GetAllHrefsFromFile(htmlFilePath string) ([]string, error) {
	var result []string

	// get html.Node
	node, errN := getHTMLNodeFromFile(htmlFilePath)
	if errN != nil {
		return nil, fmt.Errorf("failed to get html.Node from file(%s):\n\t%v", errN)
	}

	func(n *html.Node)

	return result, nil
}
