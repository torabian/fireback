//go:build !wasm

package fireback

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
)

// GetCombinedConfigListCli builds the "config list" subcommand - the "show to user"
// half of the config-injection mechanism (GetCombinedConfigInfo does the actual
// collecting). Needs xapp (to reach every registered module's ConfigProvider), so
// unlike ConfigCommand's own per-field get/set subcommands (which only ever needed
// fireback's own package-level config var and are built once at package-init time),
// this has to be built fresh wherever xapp is available - see
// GetCommonWebServerCliActions.
func GetCombinedConfigListCli(xapp *application.Application) *cli.Command {
	return &cli.Command{
		Name:  "list",
		Usage: "Lists the combined configuration collected from fireback core plus every registered module (see ConfigProvider on application.ModuleProvider)",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "reveal-secrets",
				Usage: "Show real values for fields that look like secrets (password/token/secret/private) instead of masking them with ******",
			},
			&cli.StringFlag{
				Name:  "module",
				Usage: "Only show config from this module (matches ModuleProvider.Name, or \"fireback\" for the core config)",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			items := GetCombinedConfigInfo(xapp, c.Bool("reveal-secrets"))

			if module := c.String("module"); module != "" {
				filtered := make([]ConfigItemInfo, 0, len(items))
				for _, item := range items {
					if item.Module == module {
						filtered = append(filtered, item)
					}
				}
				items = filtered
			}

			str, err := json.MarshalIndent(items, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(str))

			return nil
		},
	}
}
