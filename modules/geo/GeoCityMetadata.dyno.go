package geo

import (
	reflect "reflect"

	"github.com/torabian/fireback/modules/workspaces"
)

var GeoCityEntityMetaConfig map[string]int64 = map[string]int64{}

var GeoCityEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoCityEntity{}))
