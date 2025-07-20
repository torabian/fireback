package fireback

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
	"gopkg.in/yaml.v2"
)

func HandleXsrc[T any](c *cli.Context, template *T) {
	if c.IsSet("x-src") {
		path := c.String("x-src")
		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file: %v\n", err)
			os.Exit(1)
		}

		ext := strings.ToLower(filepath.Ext(path))
		switch ext {
		case ".yaml", ".yml":
			if err := yaml.Unmarshal(data, template); err != nil {
				fmt.Fprintf(os.Stderr, "YAML parse error: %v\n", err)
				os.Exit(1)
			}
		case ".json":
			if err := json.Unmarshal(data, template); err != nil {
				fmt.Fprintf(os.Stderr, "JSON parse error: %v\n", err)
				os.Exit(1)
			}
		default:
			// Unknown extension: try YAML first, then JSON
			if err := yaml.Unmarshal(data, template); err != nil {
				if err := json.Unmarshal(data, template); err != nil {
					fmt.Fprintf(os.Stderr, "Unrecognized format: %v\n", err)
					os.Exit(1)
				}
			}
		}
	}
}
