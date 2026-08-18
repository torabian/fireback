// firebackModuleCliHandlers - split out of FirebackModule.go (which is otherwise
// fine under wasm - just names the module and points at its migrations dir) since
// PushNotificationCmd (VapidCli.go) is !wasm-only. See FirebackModuleCliWasm.go
// for wasm's own (empty) counterpart.
//go:build !wasm

package fireback

import "github.com/urfave/cli/v3"

func firebackModuleCliHandlers() []*cli.Command {
	return []*cli.Command{
		&PushNotificationCmd,
	}
}
