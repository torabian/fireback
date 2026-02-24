package fireback

import (
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/emi/lib/core"
	"github.com/urfave/cli"
)

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

// This is a helper, which would convert the emi actions created,
// and cast them into the urfave 2/3.
func CastEmiFlagToUrfave(flags []emigo.CliFlag) []cli.Flag {
	var out []cli.Flag

	for _, f := range flags {
		// Recursively flatten children
		if len(f.Children) > 0 {
			out = append(out, CastEmiFlagToUrfave(f.Children)...)
			continue
		}

		usage := firstNonEmpty(f.Description, f.Usage)
		req := f.Required

		switch {
		case strings.Contains(f.Type, "int") && !core.IsNullable(f.Type):
			out = append(out, &cli.Int64Flag{Name: f.Name, Usage: usage, Required: req})
		case strings.Contains(f.Type, "bool") && !core.IsNullable(f.Type):
			out = append(out, &cli.BoolFlag{Name: f.Name, Usage: usage, Required: req})
		case strings.Contains(f.Type, "float") && !core.IsNullable(f.Type):
			out = append(out, &cli.Float64Flag{Name: f.Name, Usage: usage, Required: req})
		default:
			out = append(out, &cli.StringFlag{Name: f.Name, Usage: usage, Required: req})
		}
	}

	return out
}
