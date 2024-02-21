package worldtimezone

import (
	reflect "reflect"

	"pixelplux.com/fireback/modules/workspaces"
)

var TimezoneGroupEntityMetaConfig map[string]int64 = map[string]int64{}

var TimezoneGroupEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&TimezoneGroupEntity{}))
