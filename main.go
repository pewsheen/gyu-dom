package main

import (
	"fmt"
	"log"
	"strings"

	"go.bytecodealliance.org/cm"
	"golang.org/x/net/html"
	"pews.dev/internal/pews/dom/domparser"
)

func init() {
	// domparser.Exports.Domparser.Constructor = func() (result domparser.Domparser) {
	// 	result = domparser.DomparserResourceNew(0)
	// 	return
	// }
	domparser.Exports.Domparser.ParseFromString = func(self cm.Rep, inputString string, mimeType string) (result domparser.Document) {
		doc, err := html.Parse(strings.NewReader(inputString))
		if err != nil {
			log.Fatal(err)
		}
		for n := range doc.Descendants() {
			fmt.Printf("Node type: %d\n", n.Type)
			fmt.Printf("Node attributes: %s\n", n.Attr)
			fmt.Printf("Node tag: %s\n", n.Data) // tag name
		}

		return
	}
}

// main is required for the `wasi` target, even if it isn't used.
func main() {}
