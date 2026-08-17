package interfacetoolsdefs

import (
	"encoding/json"
	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/abac/abaccomplexes"
)

// The base class definition for tableViewSizingDto
type TableViewSizingDto struct {
	UniqueId  emigo.Nullable[string] `json:"uniqueId" yaml:"uniqueId"`
	TableName string                 `json:"tableName" validate:"required" yaml:"tableName"`
	Sizes     string                 `json:"sizes" yaml:"sizes"`
	// The unique-id of the workspace which content belongs to.
	WorkspaceId emigo.Nullable[string] `json:"workspaceId" yaml:"workspaceId"`
	// The unique-id of the user which created/owns the record.
	UserId emigo.Nullable[string] `json:"userId" yaml:"userId"`
	// The time the record was created. Populated automatically by gorm.
	CreatedAt abaccomplexes.PlainTime `json:"createdAt" yaml:"createdAt"`
	// The time the record was last updated. Populated automatically by gorm.
	UpdatedAt abaccomplexes.PlainTime `json:"updatedAt" yaml:"updatedAt"`
}

func (x *TableViewSizingDto) Json() string {
	if x != nil {
		str, _ := json.MarshalIndent(x, "", "  ")
		return string(str)
	}
	return ""
}
