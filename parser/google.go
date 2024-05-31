package parser

import (
	"bytes"
	"log"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
)

type GoogleResponseParser struct {
	SearchElementString string
	ParsedString        string
}

func (parser *GoogleResponseParser) ParseRawElementString() error {
	// Parse the HTML of the #rso element
	doc, err := html.Parse(bytes.NewReader([]byte(parser.SearchElementString)))
	if err != nil {
		log.Fatalf("Failed to parse the HTML: %v", err)
	}

	// Buffer to store the cleaned text
	var buffer bytes.Buffer

	// Traverse the HTML nodes
	traverse(doc, &buffer)

	// Sanitize the text
	p := bluemonday.StrictPolicy()
	cleanedText := p.SanitizeBytes(buffer.Bytes())
	parser.ParsedString = string(cleanedText)

	return nil
}

// Function to traverse the HTML nodes and remove tags but keep text
func traverse(n *html.Node, buffer *bytes.Buffer) {
	if n.Type == html.TextNode {
		buffer.WriteString(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Skip script and style tags and their content
		if c.Type == html.ElementNode && (c.Data == "script" || c.Data == "style") {
			continue
		}

		traverse(c, buffer)
		if c.Type == html.ElementNode && c.Data == "div" {
			if buffer.Len() < 2 || buffer.String()[buffer.Len()-2:] != "\n\n" {
				buffer.WriteString("\n")
			}
		}
	}
}
