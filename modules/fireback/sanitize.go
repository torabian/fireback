package fireback

import "github.com/microcosm-cc/bluemonday"

var StripPolicy = bluemonday.StripTagsPolicy()
var UgcPolicy = bluemonday.UGCPolicy().AllowAttrs("class").Globally()
