package fireback

func GetCapabilitiesAction(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error) {
	query, err := ResolveActionContext(c, &SecurityModel{
		AllowOnRoot:    true,
		ActionRequires: []PermissionInfo{PERM_ROOT_CAPABILITY_QUERY},
	})

	if err != nil {
		return nil, err
	}

	res, qrm, err2 := CapabilityActions.Query(*query)
	if err2 != nil {
		return nil, err
	}

	return &GetCapabilitiesActionResponse{
		Payload: GResponseQuery(res, qrm, query),
	}, nil
}

// func GetCapabilitiesActionCliHandler(
// 	handler func(c GetCapabilitiesActionRequest) (*GetCapabilitiesActionResponse, error),
// ) *cli.Command {
// 	meta := GetCapabilitiesActionMeta()

// 	return &cli.Command{
// 		Name:        meta.Name,
// 		Description: meta.Description,
// 		Action: func(ctx context.Context, c *cli.Command) error {
// 			req := GetCapabilitiesActionRequest{
// 				CliCtx: c,
// 			}

// 			res, err := handler(req)
// 			if err != nil {
// 				return err
// 			}

// 			v, err := json.MarshalIndent(res, "", "  ")
// 			if err != nil {
// 				return err
// 			}

// 			fmt.Print(string(v), err)
// 			return nil
// 		},
// 	}
// }
