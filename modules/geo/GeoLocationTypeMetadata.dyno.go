package geo

import (
	reflect "reflect"

	"github.com/torabian/fireback/modules/workspaces"
)

var GeoLocationTypeEntityMetaConfig map[string]int64 = map[string]int64{}

var GeoLocationTypeEntityJsonSchema = workspaces.ExtractEntityFields(reflect.ValueOf(&GeoLocationTypeEntity{}))
