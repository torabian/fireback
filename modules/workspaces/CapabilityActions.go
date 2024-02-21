package workspaces

import (
	"fmt"
	"sort"
	"strings"
)

func CapabilityActionCreate(
	dto *CapabilityEntity, query QueryDSL,
) (*CapabilityEntity, *IError) {
	return CapabilityActionCreateFn(dto, query)
}

func CapabilityActionUpdate(
	query QueryDSL,
	fields *CapabilityEntity,
) (*CapabilityEntity, *IError) {
	return CapabilityActionUpdateFn(query, fields)
}

type CapabilityChild struct {
	UniqueId   string             `protobuf:"bytes,1,opt,name=uniqueId,proto3" json:"uniqueId,omitempty"`
	Children   []*CapabilityChild `protobuf:"bytes,2,rep,name=children,proto3" json:"children,omitempty"`
	Visibility *string            `protobuf:"bytes,3,opt,name=visibility,proto3,oneof" json:"visibility,omitempty" yaml:"visibility"`
	Updated    int64              `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created    int64              `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty" gorm:"autoUpdateTime:nano"`
}
type CapabilitiesResult struct {
	Capabilities []*CapabilityEntity `protobuf:"bytes,1,rep,name=capabilities,proto3" json:"capabilities,omitempty"`
	Nested       []*CapabilityChild  `protobuf:"bytes,2,rep,name=nested,proto3" json:"nested,omitempty"`
	Visibility   *string             `protobuf:"bytes,3,opt,name=visibility,proto3,oneof" json:"visibility,omitempty" yaml:"visibility"`
	Updated      int64               `protobuf:"varint,4,opt,name=updated,proto3" json:"updated,omitempty" gorm:"autoUpdateTime:nano"`
	Created      int64               `protobuf:"varint,5,opt,name=created,proto3" json:"created,omitempty" gorm:"autoUpdateTime:nano"`
}

func treeToCapabilityChild(items []NestedNode) []*CapabilityChild {
	data := []*CapabilityChild{}

	for _, item := range items {

		children := []*CapabilityChild{}

		if len(item.Children) > 0 {
			children = treeToCapabilityChild(item.Children)
		}

		data = append(data, &CapabilityChild{
			UniqueId: item.UniqueId,
			Children: children,
		})
	}

	return data
}

func CapabilityActionGetTree(query QueryDSL) (*CapabilitiesResult, *IError) {

	// Read the comments inside CapabilityActionQuery
	items, _, err := CapabilityActionQuery(query)

	fmt.Println(items[0])
	sort.Slice(items, func(i, j int) bool {
		return items[i].UniqueId < items[j].UniqueId
	})

	tree := Tree{}
	for _, item := range items {
		if item.UniqueId == "" {
			continue
		}
		if strings.HasSuffix(item.UniqueId, "/*") {
			tree.Add(strings.TrimRight(item.UniqueId, "/*"))
		} else {
			tree.Add(item.UniqueId)
		}
	}
	itemsa := tree.ToObject(true)

	return &CapabilitiesResult{
		Capabilities: items,
		Nested:       treeToCapabilityChild(itemsa),
	}, GormErrorToIError(err)
}
