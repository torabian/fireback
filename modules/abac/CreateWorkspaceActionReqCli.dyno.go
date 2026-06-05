//go:build !wasm

package abac

func GetCreateWorkspaceActionReqCliFlags(prefix string) []emigo.CliFlag {
	return []emigo.CliFlag{
		{
			Name: prefix + "name",
			Type: "string",
		},
	}
}
func CastCreateWorkspaceActionReqFromCli(c emigo.CliCastable) CreateWorkspaceActionReq {
	data := CreateWorkspaceActionReq{}
	if c.IsSet("name") {
		data.Name = c.String("name")
	}
	return data
}
