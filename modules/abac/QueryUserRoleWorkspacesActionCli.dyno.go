//go:build !wasm

package abac

import (
	"github.com/torabian/emi/emigo"
	"github.com/urfave/cli/v3"
	"reflect"
)

func GetQueryUserRoleWorkspacesActionPathParameterCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name:     prefix + "pp-ms",
			Type:     "string",
			Required: true,
		},
	}
}

// Extracts the path parameter from a urfave v3 cli.
func QueryUserRoleWorkspacesActionPathParameterFromCli(c *cli.Command) QueryUserRoleWorkspacesActionPathParameter {
	return QueryUserRoleWorkspacesActionPathParameterFromFn(func(key string) string {
		// In cli, they are prefixed with pp, to avoid conflict with other params coming from 'in'
		// section of the definition.
		return c.String("pp-" + key)
	})
}
func (x QueryUserRoleWorkspacesActionRequest) IsCli() bool {
	if x.CliCtx == nil {
		return false
	}
	v := reflect.ValueOf(x.CliCtx)
	switch v.Kind() {
	case reflect.Ptr, reflect.Map, reflect.Slice, reflect.Interface, reflect.Func, reflect.Chan:
		return !v.IsNil()
	}
	return true
}
