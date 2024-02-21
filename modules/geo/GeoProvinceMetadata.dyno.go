package geo

import (
	reflect "reflect"

	"github.com/torabian/fireback/modules/workspaces"
)

var GeoProvinceEntityMetaConfig map[string]int64 = map[string]int64{}

var GeoProvinceEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoProvinceEntity{}))
