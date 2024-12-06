package htmlparser

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Get raw HTML data from .html file
func GetHTMLNodeFromFile(htmlFilePath string) (*html.Node, error) {
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

// Process node
func processNodeByTag(n *html.Node) (string, error) {
	if n.FirstChild != nil {
		fmt.Println("Name:", n.FirstChild.Data)
		return n.FirstChild.Data, nil
	}

	return "", fmt.Errorf("???")
}

// Get all hrefs of <a> of html file
func GetAllHrefsFromFile(n *html.Node) ([]string, error) {
	var result []string

	var processAll func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			processNodeByTag(n)
		}
	} 

	return result, nil
}
