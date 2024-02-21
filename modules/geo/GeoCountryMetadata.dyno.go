package geo

import (
	reflect "reflect"

	"github.com/torabian/fireback/modules/workspaces"
)

var GeoCountryEntityMetaConfig map[string]int64 = map[string]int64{}

var GeoCountryEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoCountryEntity{}))
