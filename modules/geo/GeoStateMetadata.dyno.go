package geo

import (
	reflect "reflect"

	"pixelplux.com/fireback/modules/workspaces"
)

var GeoStateEntityMetaConfig map[string]int64 = map[string]int64{}

var GeoStateEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoStateEntity{}))
