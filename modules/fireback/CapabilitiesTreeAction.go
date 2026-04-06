package fireback

import (
	"sort"
	"strings"
)

func init() {
	// Override the implementation with our actual code.
	CapabilitiesTreeImpl = CapabilitiesTreeAction
}

func CapabilitiesTreeAction(c CapabilitiesTreeActionRequest, query QueryDSL) (*CapabilitiesTreeActionResponse, error) {

	// Read the comments inside CapabilityActionQuery
	query.ItemsPerPage = 9999
	items, _, err := CapabilityActions.Query(query)
	if err != nil {
		return nil, GormErrorToIError(err)
	}

	itemsFiltered := []CapabilityInfoDto{}

	workspaceAccesses, rolesPermission := GetWorkspaceAndUserAccesses(query)
	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	tree := Tree{}

	for _, item := range items {
		if item == nil || item.UniqueId == "" {
			continue
		}

		// Filter based on the workspace and role and not allow to create more access than the user has.
		meetsUser := MeetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, rolesPermission)
		meetsWorkspace := MeetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, workspaceAccesses)

		if !meetsUser || !meetsWorkspace {
			continue
		}

		itemsFiltered = append(itemsFiltered, CapabilityInfoDto{
			UniqueId: item.UniqueId,
			Name:     item.Name,
		})

		if strings.HasSuffix(item.UniqueId, ".*") {
			tree.Add(strings.TrimRight(item.UniqueId, ".*"), ".")
		} else {
			tree.Add(item.UniqueId, ".")
		}
	}

	itemsa := tree.ToObject(true)

	return &CapabilitiesTreeActionResponse{
		Payload: GResponseSingleItem(CapabilitiesTreeActionRes{
			Capabilities: itemsFiltered,
			Nested:       treeToCapabilityChild(itemsa),
		}),
	}, nil

}

func treeToCapabilityChild(items []NestedNode) []CapabilityInfoDto {
	data := []CapabilityInfoDto{}

	for _, item := range items {

		children := []CapabilityInfoDto{}

		if len(item.Children) > 0 {
			children = treeToCapabilityChild(item.Children)
		}

		data = append(data, CapabilityInfoDto{
			UniqueId: item.UniqueId,
			Children: children,
		})
	}

	return data
}
