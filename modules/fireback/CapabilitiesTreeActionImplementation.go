package fireback

import (
	"fmt"
	"sort"
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback/terminal"
)

func CapabilitiesTreeAction(c CapabilitiesTreeActionRequest) (*CapabilitiesTreeActionResponse, error) {

	query, err := ResolveActionContext(c, &SecurityModel{
		ResolveStrategy: ResolveStrategyUser,
		ActionRequires: []PermissionInfo{
			PERM_ROOT_CAPABILITY_QUERY,
		},
	})
	if err != nil {
		return nil, err
	}

	query.ItemsPerPage = 9999
	query.WorkspaceId = "system"
	query.UserId = "system"

	var items []*CapabilityEntity

	err2 := GetDbRef().Debug().Model(&CapabilityEntity{}).Limit(1000).Find(&items).Error
	if err2 != nil {
		return nil, GormErrorToIError(err2)
	}

	itemsFiltered := []CapabilityInfoDto{}

	workspaceAccesses, rolesPermission := GetWorkspaceAndUserAccesses(*query)
	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	fmt.Println(4, workspaceAccesses, rolesPermission)

	tree := terminal.Tree{}

	for _, item := range items {
		if item == nil || item.UniqueId == "" {
			continue
		}

		// Filter based on the workspace and role and not allow to create more access than the user has.
		meetsUser := MeetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, rolesPermission)
		meetsWorkspace := MeetsCheck([]PermissionInfo{{CompleteKey: item.UniqueId}}, workspaceAccesses)

		fmt.Println("1", meetsUser, meetsWorkspace, item.UniqueId, rolesPermission, workspaceAccesses)
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
			Capabilities: emigo.CollectionReplace(itemsFiltered),
			Nested:       emigo.CollectionReplace(treeToCapabilityChild(itemsa)),
		}),
	}, nil

}

func treeToCapabilityChild(items []terminal.NestedNode) []CapabilityInfoDto {
	data := []CapabilityInfoDto{}

	for _, item := range items {

		children := []CapabilityInfoDto{}

		if len(item.Children) > 0 {
			children = treeToCapabilityChild(item.Children)
		}

		data = append(data, CapabilityInfoDto{
			UniqueId: item.UniqueId,
			Children: emigo.CollectionReplace(children),
		})
	}

	return data
}
