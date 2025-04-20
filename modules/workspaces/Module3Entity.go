package workspaces

// Default events for every entity that we generate.
// Maybe allow developers to customise it later
func (x *Module3Entity) DefaultEvents() []*Module3Event {

	allowOnRoot := false
	resolveStrategy := ""
	if x.SecurityModel != nil && x.SecurityModel.ReadOnRoot != nil {
		allowOnRoot = *x.SecurityModel.ReadOnRoot
	}
	if x.SecurityModel != nil && x.SecurityModel.ResolveStrategy != nil {
		resolveStrategy = *x.SecurityModel.ResolveStrategy
	}

	cacheKey := x.EntityName()
	if x.RootModule != nil {
		cacheKey = "*" + x.RootModule.Name + "." + cacheKey
	}

	return []*Module3Event{
		{
			Name: x.Upper() + "Created",
			Payload: &Module3ActionBody{
				Entity: x.Upper() + "Entity",
			},
			CacheKey: cacheKey,
			SecurityModel: &SecurityModel{
				AllowOnRoot:     allowOnRoot,
				ResolveStrategy: resolveStrategy,
				ActionRequires: []PermissionInfo{
					{
						CompleteKey: "PERM_ROOT_" + x.AllUpper() + "_QUERY",
					},
				},
			},
		},
		{
			Name: x.Upper() + "Updated",
			Payload: &Module3ActionBody{
				Entity: x.Upper() + "Entity",
			},
			CacheKey: cacheKey,
			SecurityModel: &SecurityModel{
				AllowOnRoot:     allowOnRoot,
				ResolveStrategy: resolveStrategy,
				ActionRequires: []PermissionInfo{
					{
						CompleteKey: "PERM_ROOT_" + x.AllUpper() + "_QUERY",
					},
				},
			},
		},
	}

}
