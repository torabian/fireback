// PopulateInteractively is CLI-only (walks flags, prompting AskForInputOptional/
// AskBoolean - both package-level func vars only ever assigned by
// clitools/Prompts.go, itself !wasm) - split out of Query.go so the rest of that
// file's genuinely wasm-needed query-building logic doesn't drag CliInteractiveFlag
// (defined in the also-!wasm CliActions.go) along with it.
//go:build !wasm

package fireback

import "github.com/urfave/cli/v3"

func PopulateInteractively[T any](entity T, c *cli.Command, flags []CliInteractiveFlag) {
	for _, item := range flags {
		if (!item.Required && !item.Recommended) && !c.Bool("all") {
			continue
		}

		if item.Type == "string" {
			var result string
			if !item.Required {
				result, _, _ = AskForInputOptional(item.Name, "")
			} else {
				result, _, _ = AskForInputOptional(item.Name, "")
			}
			SetField(entity, ToLower(item.StructField), &result)
		}
		if item.Type == "bool" {
			result := AskBoolean(item.Name)
			SetField(entity, ToLower(item.StructField), &result)
		}

	}
}
