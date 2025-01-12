package workspaces

import (
	"slices"
)

var AVAILABLE_FIREBACK_DATA_TYPES = []string{
	FIELD_TYPE_ARRAY,
	FIELD_TYPE_ARRAYP,
	FIELD_TYPE_JSON,
	FIELD_TYPE_ONE,
	FIELD_TYPE_DATE,
	FIELD_TYPE_MANY2MANY,
	FIELD_TYPE_OBJECT,
	FIELD_TYPE_EMBED,
	FIELD_TYPE_ENUM,
	FIELD_TYPE_COMPUTED,
	FIELD_TYPE_TEXT,
	FIELD_TYPE_STRING,
	FIELD_TYPE_ANY,
}

func (x *Module3Field) DialectValidate() []*IErrorItem {
	res := []*IErrorItem{}

	if !slices.Contains(AVAILABLE_FIREBACK_DATA_TYPES, x.Type) {
		res = append(res, &IErrorItem{
			// Message: &WorkspacesMessages.DataTypeDoesNotExistsInFireback,
		})
	}

	return res
}
