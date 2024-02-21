package geo

import (
	reflect "reflect"

	"github.com/torabian/fireback/modules/workspaces"
)

var LocationDataEntityMetaConfig map[string]int64 = map[string]int64{}

var LocationDataEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&LocationDataEntity{}))
