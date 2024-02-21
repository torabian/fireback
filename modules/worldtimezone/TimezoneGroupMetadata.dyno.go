package worldtimezone

import (
	reflect "reflect"

	"github.com/torabian/fireback/modules/workspaces"
)

var TimezoneGroupEntityMetaConfig map[string]int64 = map[string]int64{}

var TimezoneGroupEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&TimezoneGroupEntity{}))
