package fireback

import (
	"sort"
	"strings"
)

func init() {
	// Override the implementation with our actual code.
	CapabilitiesTreeActionImp = CapabilitiesTreeAction
}

func CapabilitiesTreeAction(query QueryDSL) (*CapabilitiesTreeActionResDto, *IError) {

	// Read the comments inside CapabilityActionQuery
	items, _, err := CapabilityActions.Query(query)
	itemsFiltered := []*CapabilityEntity{}

	workspaceAccesses, rolesPermission := GetWorkspaceAndUserAccesses(query)
	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	tree := Tree{}

	for _, item := range items {
		if item.UniqueId == "" {
			continue
		}

		// Filter based on the workspace and role and not allow to create more access than the user has.
		meetsUser := MeetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, rolesPermission)
		meetsWorkspace := MeetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, workspaceAccesses)

		if !meetsUser || !meetsWorkspace {
			continue
		}

		itemsFiltered = append(itemsFiltered, item)
		if strings.HasSuffix(item.UniqueId, ".*") {
			tree.Add(strings.TrimRight(item.UniqueId, ".*"), ".")
		} else {
			tree.Add(item.UniqueId, ".")
		}
	}
	itemsa := tree.ToObject(true)

	return &CapabilitiesTreeActionResDto{
		Capabilities: itemsFiltered,
		Nested:       treeToCapabilityChild(itemsa),
	}, GormErrorToIError(err)

}

func treeToCapabilityChild(items []NestedNode) []*CapabilityEntity {
	data := []*CapabilityEntity{}

	for _, item := range items {

		children := []*CapabilityEntity{}

		if len(item.Children) > 0 {
			children = treeToCapabilityChild(item.Children)
		}

		data = append(data, &CapabilityEntity{
			UniqueId: item.UniqueId,
			Children: children,
		})
	}

	return data
}
