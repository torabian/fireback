package fireback

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var IntelisenseTest = []Test{
	{
		Name: "Test for correct detection of the place in intelisense",
		Function: func(t *TestContext) error {
			file, _ := os.ReadFile("modules/fireback/WorkspaceModule3.yml")
			// Parse the YAML data into a Node

			fmt.Println(GetContextFromYaml((file), 1, 1))

			return nil
		},
	},
}

type YamlData struct {
	Line int
	Char int
	Path string
}

func GetContextFromYaml(content []byte, line int, col int) string {
	var rootNode yaml.Node
	err := yaml.Unmarshal(content, &rootNode)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var items = []YamlData{
		{
			Char: 0,
			Line: 0,
			Path: "root",
		},
	}

	items = append(items, iterateYAML3(&rootNode, "")...)

	var suggest *YamlData = nil
	for i := 0; i < len(items)-1; i++ {
		item := items[i]
		// next := items[i+1]
		// For debugging
		// fmt.Println(item.Line, item.Char, "looking", line, col, item.Path)

		if item.Line == line {
			if item.Char < col {
				suggest = &item
			}
		}

	}

	if suggest != nil {
		return suggest.Path
	}

	for i := len(items) - 1; i > 0; i-- {
		item := items[i]

		if item.Line < line && item.Char < col {
			suggest = &item
			break
		}

	}

	if suggest != nil {
		return suggest.Path
	}

	return "root"
}

func iterateYAML3(node *yaml.Node, indent string) []YamlData {
	var run func(node *yaml.Node, indent string, prefix string) []YamlData
	run = func(node *yaml.Node, indent string, prefix string) []YamlData {
		items := []YamlData{}
		switch node.Kind {
		case yaml.DocumentNode:
			for _, content := range node.Content {

				items = append(items, run(content, indent, "root")...)
			}
		case yaml.MappingNode:
			for i := 0; i < len(node.Content); i += 2 {
				keyNode := node.Content[i]
				valueNode := node.Content[i+1]

				item := YamlData{
					Path: prefix + "/" + keyNode.Value,
					Line: keyNode.Line,
					Char: keyNode.Column,
				}
				items = append(items, item)
				items = append(items, run(valueNode, indent+"  ", prefix+"/"+keyNode.Value)...)

			}
		case yaml.SequenceNode:
			for i, content := range node.Content {

				item := YamlData{
					Path: prefix + fmt.Sprintf("/%d", i),
					Line: content.Line,
					Char: content.Column,
				}

				items = append(items, item)
				items = append(items, run(content, indent+"  ", prefix+fmt.Sprintf("/%d", i))...)
			}
		case yaml.ScalarNode:
			items = append(items, YamlData{
				Path: prefix + "/value?actual=" + node.Value,
				Line: node.Line,
				Char: node.Column,
			})
		}

		return items

	}

	return run(node, indent, "")
}
