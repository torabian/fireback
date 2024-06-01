package workspaces

import (
	"fmt"
	"log"
	"net/url"

	"github.com/sourcegraph/go-lsp"
	"gopkg.in/yaml.v3"
)

// ParseYAML parses the YAML file into a tree of nodes.
func ParseYAML(data []byte) (*yaml.Node, error) {
	var root yaml.Node
	err := yaml.Unmarshal(data, &root)
	if err != nil {
		return nil, err
	}
	return &root, nil
}

type URLParams map[string]string

// FindContext traverses the YAML node tree to find the context based on line and character number.

var result string

func PrintNode(node *yaml.Node, indent string, p string, com *string, line int, offset int) {

	log.Println(node.Kind, line)
	switch node.Kind {
	case yaml.DocumentNode:
		for _, content := range node.Content {
			PrintNode(content, indent+"  ", p, com, line, offset)
		}
	case yaml.MappingNode:

		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			value := node.Content[i+1]
			cline := node.Content[i].Line
			newP := p + "/" + key.Value
			if cline == line {
				result = newP
				*com = result
			}

			PrintNode(value, indent+"    ", newP, com, line, offset)
		}
	case yaml.SequenceNode:
		log.Println("Sequence catched", p, com)
		for i, item := range node.Content {
			cline := node.Content[i].Line
			log.Println("Sequence:", i, cline, line)
			newP := p + fmt.Sprintf("/%v", i)
			if cline == line {
				result = newP
				*com = result
			}

			PrintNode(item, indent+"    ", newP, com, line, offset)
		}
	case yaml.ScalarNode:
		if node.Line == line {
			result = p + "?value=" + node.Value
			*com = result
		}
	}
}

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

func GetLineContext(yamlData []byte, line int, offset int) string {

	rootNode, err := ParseYAML(yamlData)

	// fmt.Println(rootNode)
	if err != nil {
		log.Println(err)
		return ""
	}

	// Example line and character number.
	// line := 3
	// char := 5

	var p string = ""
	var comp string

	log.Println("Line:", line, "offset", offset)
	PrintNode(rootNode, "  ", p, &comp, line, offset)

	return comp
}
