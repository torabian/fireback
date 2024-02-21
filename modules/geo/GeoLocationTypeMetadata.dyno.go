package geo

import (
	reflect "reflect"

	"pixelplux.com/fireback/modules/workspaces"
)

var GeoLocationTypeEntityMetaConfig map[string]int64 = map[string]int64{}

var GeoLocationTypeEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoLocationTypeEntity{}))
