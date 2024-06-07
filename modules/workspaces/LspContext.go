package workspaces

import (
	"fmt"
	"net/url"

	"github.com/sourcegraph/go-lsp"
)

// ParseYAML parses the YAML file into a tree of nodes.
// func ParseYAML(data []byte) (*yaml.Node, error) {
// 	var root yaml.Node
// 	err := yaml.Unmarshal(data, &root)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &root, nil
// }

type URLParams map[string]string

// FindContext traverses the YAML node tree to find the context based on line and character number.

var result string

func removeQueryParams(rawURL string) (string, error) {
	// Parse the URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %v", err)
	}

	// Clear the query parameters
	parsedURL.RawQuery = ""

	// Reconstruct the URL without query parameters
	return parsedURL.String(), nil
}

type LspHandler struct {
	Path    string
	Handler func() []lsp.CompletionItem
}
