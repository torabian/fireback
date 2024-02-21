package geo

import (
	reflect "reflect"

	"pixelplux.com/fireback/modules/workspaces"
)

var GeoLocationEntityMetaConfig map[string]int64 = map[string]int64{}

var GeoLocationEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoLocationEntity{}))
