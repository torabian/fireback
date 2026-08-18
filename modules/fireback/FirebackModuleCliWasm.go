// See FirebackModuleCli.go's own doc comment - wasm's counterpart just has
// nothing to add (no VAPID push notification CLI command under wasm).
//go:build wasm

package fireback

import "github.com/urfave/cli/v3"

func firebackModuleCliHandlers() []*cli.Command {
	return nil
}
