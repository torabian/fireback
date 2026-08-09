package abac

import (
	"sort"
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/torabian/fireback/modules/fireback/terminal"
)

func CapabilitiesTreeAction(c CapabilitiesTreeActionRequest) (*CapabilitiesTreeActionResponse, error) {

	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{
		ResolveStrategy: fireback.ResolveStrategyUser,
		ActionRequires: []application.PermissionInfo{
			PERM_ROOT_CAPABILITY_QUERY,
		},
	})

	if err != nil {
		return nil, err
	}

	query.ItemsPerPage = 9999

	var items []*CapabilityEntity

	err2 := fireback.GetDbRef().Debug().Model(&CapabilityEntity{}).Limit(1000).Find(&items).Error
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}

	itemsFiltered := []CapabilityInfoDto{}

	workspaceAccesses, rolesPermission := GetWorkspaceAndUserAccesses(*query)

	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	tree := terminal.Tree{}

	for _, item := range items {
		if item == nil || item.UniqueId == "" {
			continue
		}

		// Filter based on the workspace and role and not allow to create more access than the user has.
		meetsUser := MeetsCheck([]application.PermissionInfo{{CompleteKey: item.UniqueId}}, rolesPermission)
		meetsWorkspace := MeetsCheck([]application.PermissionInfo{{CompleteKey: item.UniqueId}}, workspaceAccesses)

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
		Payload: fireback.GResponseSingleItem(CapabilitiesTreeActionRes{
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
